package scores

import (
	"database/sql"
	"fmt"
	"github.com/Swan/Nameless/src/auth"
	"github.com/Swan/Nameless/src/common"
	"github.com/Swan/Nameless/src/db"
	"github.com/Swan/Nameless/src/db/achievements"
	"github.com/Swan/Nameless/src/handlers"
	"github.com/Swan/Nameless/src/processors"
	"github.com/Swan/Nameless/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Handler struct {
	scoreData       scoreSubmissionData
	user            db.User
	mapData         db.Map
	mapPath         string
	stats           db.UserStats
	difficulty      processors.DifficultyProcessor
	rating          processors.RatingProcessor
	oldPersonalBest db.Score
	newScoreId      int64
	unlockedAchievements []achievements.Achievement
}

func (h Handler) SubmitPOST(c *gin.Context) {
	timeStart := time.Now()
	
	var err error
	h.user, err = auth.GetInGameUser(c)

	if err != nil {
		log.Errorf("Could not authenticate user - %v\n", err)
		handlers.ReturnError(c, http.StatusForbidden, err.Error())
		return
	}

	h.scoreData, err = parseScoreSubmissionData(c)

	if err != nil {
		h.logIgnoringScore(fmt.Sprintf("Invalid score data - %v", err))
		handlers.Return400(c)
		return
	}

	hasRankedMods := common.IsModComboRanked(h.scoreData.Mods)

	if !hasRankedMods {
		h.logIgnoringScore(fmt.Sprintf("unranked modifiers -%v", common.GetModsString(h.scoreData.Mods)))
		handlers.Return400(c)
		return
	}

	h.mapData, err = db.GetMapByMD5(h.scoreData.MapMD5)

	if err != nil {
		if err == sql.ErrNoRows {
			h.logIgnoringScore(fmt.Sprintf("Unknown Map - `%v`", h.scoreData.MapMD5))
			handlers.Return400(c)
			return 
		}
		
		h.logError(fmt.Sprintf("Failure fetching map `%v` - %v", h.scoreData.MapMD5, err))
		handlers.Return500(c)
		return
	}

	err = h.scoreData.validateGameMode(&h.mapData)

	if err != nil {
		h.logIgnoringScore(fmt.Sprintf("non-matching game modes - %v", err))
		handlers.Return400(c)
		return
	}

	h.mapPath, err = utils.CacheQuaFile(h.mapData)
	
	if err != nil {
		h.logError(fmt.Sprintf("Unable to catch map file `%v` - %v", h.mapData.Id, err))
		handlers.Return500(c)
		return
	}

	h.stats, err = db.GetUserStats(h.user.Id, h.scoreData.GameMode)

	if err != nil {
		h.logError(fmt.Sprintf("Failed to fetch user stats - %v", err))
		handlers.Return500(c)
		return
	}

	err = h.handleSubmission(c)

	// Responses are given to the player inside of handleSubmission, so it's not needed here
	if err != nil {
		h.logError(fmt.Sprintf("Failed to submit score - `%v", err))
		return
	}
	
	h.sendSuccessfulResponse(c)
	h.logScore(time.Since(timeStart))
}

// Handles submitting the score into the database, achievements, leaderboards, etc
func (h *Handler) handleSubmission(c *gin.Context) error {
	err := h.checkZeroTotalScore(c)

	if err != nil {
		return err
	}

	err = h.checkDuplicateScore(c)

	if err != nil {
		return err
	}

	err = h.calculatePerformanceRating(c)

	if err != nil {
		return err
	}

	err = h.updateOldPersonalBest(c)

	if err != nil {
		return err
	}

	err = h.insertScore(c)

	if err != nil {
		return err
	}

	err = h.updateUserLatestActivity(c)

	if err != nil {
		return err
	}

	err = h.uploadReplay(c)

	if err != nil {
		return err
	}

	err = h.updateScoreboardCache(c)

	if err != nil {
		return err
	}

	err = h.updateUserTotalHits(c)

	if err != nil {
		return err
	}

	// Anything below this point requires the map to be ranked
	// since there will be updating leaderboards, handling achievements, etc.
	if h.mapData.RankedStatus != common.StatusRanked {
		return nil
	}

	err = h.updateUserStats(c)

	if err != nil {
		return err
	}

	err = h.updateLeaderboards(c)

	if err != nil {
		return err
	}

	err = h.handleFirstPlaceScore(c)

	if err != nil {
		return err
	}

	err = h.unlockAchievements(c)
	
	if err != nil {
		return err
	}
	
	h.updateElasticSearch()

	return nil
}

// Checks if the score has zero total score (no notes hit whatsoever). These scores
// are ignored because they are considered useless.
func (h *Handler) checkZeroTotalScore(c *gin.Context) error {
	if !h.scoreData.isValidTotalScore() {
		handlers.Return400(c)
		return fmt.Errorf("ignoring submitted score with 0 total score")
	}

	return nil
}

// Players can sometimes submit duplicate scores unexpectedly (ex. server restarts, timeouts, etc)
// This checks if the score is a duplicate, and will return a 400 if it is.
func (h *Handler) checkDuplicateScore(c *gin.Context) error {
	s, err := db.GetScoreByReplayMD5(&h.user, h.scoreData.ReplayMD5)

	// No error returned, which means a duplicate score was found
	if err == nil {
		handlers.Return400(c)
		return fmt.Errorf("duplicate submitted score found - `#%v`\n", s.Id)
	}

	// No duplicate score - everything is OK, so nil is returned here.
	if err == sql.ErrNoRows {
		return nil
	}

	handlers.Return500(c)
	return fmt.Errorf("error while attempting to fetch duplicate score - %v\n", err)
}

/// Calculates the difficulty and performance rating of the score and sets them on the handler.
func (h *Handler) calculatePerformanceRating(c *gin.Context) error {
	var err error

	h.difficulty, err = processors.CalcDifficulty(h.mapPath, h.scoreData.Mods)

	if err != nil {
		handlers.Return500(c)
		return fmt.Errorf("error while calculating difficulty rating - %v", err)
	}

	diffVal := h.difficulty.Result.OverallDifficulty
	h.rating, err = processors.CalcPerformance(diffVal, h.scoreData.Accuracy, h.scoreData.Failed)

	return nil
}

// Fetches the old personal best score and makes it no longer a PB if it isn't
func (h *Handler) updateOldPersonalBest(c *gin.Context) error {
	var err error

	h.oldPersonalBest, err = db.GetPersonalBestScore(&h.user, &h.mapData)

	// Existing personal best score was found,
	if err == nil {
		err = h.unsetOldPersonalBest()

		if err != nil {
			return err
		}

		return nil
	}

	// No personal best found, everything is OK
	if err == sql.ErrNoRows {
		return nil
	}

	handlers.Return500(c)
	return fmt.Errorf("error while fetching old personal best - %v", err)
}

// Checks if the user has beat their old PB and updates the old PB in the database
func (h *Handler) unsetOldPersonalBest() error {
	if !h.isPersonalBestScore() {
		return nil
	}

	const query string = "UPDATE scores SET personal_best = 0 WHERE id = ?"
	_, err := db.SQL.Exec(query, h.oldPersonalBest.Id)

	if err != nil {
		return err
	}

	return nil
}

// Returns if the incoming score is a personal best
func (h *Handler) isPersonalBestScore() bool {
	return !h.scoreData.Failed && h.rating.Rating > h.oldPersonalBest.PerformanceRating
}

// Inserts the incoming score into the database
func (h *Handler) insertScore(c *gin.Context) error {
	const query string = "INSERT INTO scores VALUES " +
		"(NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	grade := common.GetGradeFromAccuracy(h.scoreData.Accuracy, h.scoreData.Failed)
	isDonorScore := h.mapData.RankedStatus != common.StatusRanked
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	result, err := db.SQL.Exec(query,
		h.user.Id, h.mapData.MD5, h.scoreData.ReplayMD5, timestamp, h.scoreData.GameMode,
		h.isPersonalBestScore(), h.rating.Rating, h.scoreData.Mods, h.scoreData.Failed,
		h.scoreData.TotalScore, h.scoreData.Accuracy, h.scoreData.MaxCombo, h.scoreData.CountMarv,
		h.scoreData.CountPerf, h.scoreData.CountGreat, h.scoreData.CountGood, h.scoreData.CountOkay,
		h.scoreData.CountMiss, grade, h.scoreData.ScrollSpeed, h.scoreData.TimePlayStart,
		h.scoreData.TimePlayEnded, utils.GetIpFromRequest(c), h.scoreData.ExecutingAssemblyMD5,
		h.scoreData.EntryAssemblyMD5, h.scoreData.ReplayVersion, h.scoreData.PauseCount, h.rating.Version,
		h.difficulty.Result.Version, isDonorScore, h.scoreData.GameId)
	
	errStr := "Failed to insert score - %v"
	
	if err != nil {
		h.logError(fmt.Sprintf(errStr, err))
		handlers.Return500(c)
		return err
	}

	h.newScoreId, err = result.LastInsertId()

	if err != nil {
		h.logError(fmt.Sprintf(errStr, err))
		handlers.Return500(c)
		return err
	}

	_, err = db.Redis.Incr(db.RedisCtx, "quaver:total_scores").Result()

	if err != nil {
		h.logError(fmt.Sprintf("Failed to increment total scores in redis - %v", err))
		handlers.Return500(c)
		return err
	}

	return nil
}

// Updates the user's latest activity in the database
func (h *Handler) updateUserLatestActivity(c *gin.Context) error {
	err := db.UpdateUserLatestActivity(h.user.Id)

	if err != nil {
		h.logError(fmt.Sprintf("Failed to update user latest activity - %v", err))
		handlers.Return500(c)
		return err
	}

	return nil
}

/// Uploads the replay to azure storage
func (h *Handler) uploadReplay(c *gin.Context) error {
	if !h.isPersonalBestScore() && h.scoreData.GameId == -1 {
		return nil
	}

	fileName := fmt.Sprintf("%v.qr", h.newScoreId)

	err := utils.AzureClient.UploadFile("replays", fileName, h.scoreData.RawReplayData)

	if err != nil {
		h.logError(fmt.Sprintf("Failed to upload replay to azure - %v", err))
		handlers.Return500(c)
		return err
	}

	return nil
}

// Updates the play + fail count of the map
func (h *Handler) updateMapPlayCount(c *gin.Context) error {
	err := db.IncrementMapPlayCount(h.mapData.Id, h.scoreData.Failed)

	if err != nil {
		h.logError(fmt.Sprintf("Failed to increment map play count - %v", err))
		handlers.Return500(c)
		return err
	}

	return nil
}

/// Updates the total hits data of the user in the database & redis
func (h *Handler) updateUserTotalHits(c *gin.Context) error {
	h.stats.TotalMarv += h.scoreData.CountMarv
	h.stats.TotalPerf += h.scoreData.CountPerf
	h.stats.TotalGreat += h.scoreData.CountGreat
	h.stats.TotalGood += h.scoreData.CountGood
	h.stats.TotalOkay += h.scoreData.CountOkay
	h.stats.TotalMiss += h.scoreData.CountMiss

	err := h.stats.UpdateDatabase()

	if err != nil {
		h.logError(fmt.Sprintf("Failed to update user stats in DB - %v", err))
		handlers.Return500(c)
		return err
	}

	// Fetch stats of the other game mode, so it can be totaled up and added to the total hits leaderboard
	var otherMode common.Mode

	switch h.stats.Mode {
	case common.ModeKeys4:
		otherMode = common.ModeKeys7
	case common.ModeKeys7:
		otherMode = common.ModeKeys4
	}

	otherModeStats, err := db.GetUserStats(h.user.Id, otherMode)

	if err != nil {
		h.logError(fmt.Sprintf("Failed to fetch user stats - %v", err))
		handlers.Return500(c)
		return err
	}

	total := h.stats.GetTotalHits() + otherModeStats.GetTotalHits()

	err = db.Redis.ZAdd(db.RedisCtx, "quaver:leaderboard:total_hits_global", &redis.Z{
		Score:  float64(total),
		Member: h.user.Id,
	}).Err()

	if err != nil {
		h.logError(fmt.Sprintf("Failed to update total hits in redis - %v", err))
		handlers.Return500(c)
		return err
	}

	return nil
}

// Performs an update of the user's statistics
func (h *Handler) updateUserStats(c *gin.Context) error {
	h.stats.TotalScore += int64(h.scoreData.TotalScore)
	h.stats.PlayCount++

	if h.scoreData.Failed {
		h.stats.FailCount++
	}

	if h.scoreData.MaxCombo > h.stats.MaxCombo {
		h.stats.MaxCombo = h.scoreData.MaxCombo
	}

	if h.isPersonalBestScore() {
		// Update ranked score. If beating old personal best, take the difference between old and new score
		h.stats.RankedScore += int64(h.scoreData.TotalScore)

		if h.oldPersonalBest != (db.Score{}) {
			h.stats.RankedScore -= int64(h.oldPersonalBest.TotalScore)
		}

		// Update Overall Rating & Acc
		scores, err := db.GetUserTopScores(h.user.Id, h.scoreData.GameMode)

		if err != nil {
			h.logError(fmt.Sprintf("Failed to fetch user top scores - %v", err))
			handlers.Return500(c)
			return err
		}

		h.stats.OverallRating = db.CalculateOverallRating(scores)
		h.stats.OverallAccuracy = db.CalculateOverallAccuracy(scores)
	}

	err := h.stats.UpdateDatabase()

	if err != nil {
		h.logError(fmt.Sprintf("Failed to update user stats - %v", err))
		handlers.Return500(c)
		return err
	}

	return nil
}

// Updates the top 50 score leaderboard cache
func (h *Handler) updateScoreboardCache(c *gin.Context) error {
	if h.scoreData.Failed {
		return nil
	}

	score := h.convertToDbScore()
	err := db.UpdateScoreboardCache(&score, &h.mapData)

	if err != nil {
		h.logError(fmt.Sprintf("Failed to update scoreboard cache - %v", err))
		handlers.Return500(c)
		return err
	}

	return nil
}

// Updates the global and country leaderboards for the user.
func (h *Handler) updateLeaderboards(c *gin.Context) error {
	err := db.UpdateGlobalLeaderboard(&h.user, h.mapData.GameMode, h.stats.OverallRating)

	if err != nil {
		h.logError(fmt.Sprintf("Failed to update global leaderboard - %v", err))
		handlers.Return500(c)
		return err
	}

	err = db.UpdateCountryLeaderboard(&h.user, h.mapData.GameMode, h.stats.OverallRating)

	if err != nil {
		h.logError(fmt.Sprintf("Failed to update country leaderboard - %v", err))
		handlers.Return500(c)
		return err
	}

	return nil
}

// Checks if the user has a first place score and sends it to discord/albatross
func (h *Handler) handleFirstPlaceScore(c *gin.Context) error {
	if !h.isPersonalBestScore() {
		return nil
	}

	existingFp, err := db.GetFirstPlaceScore(h.mapData.MD5)
	gainedFirstPlace := false

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err := h.insertFirstPlaceScore()

			if err != nil {
				h.logError(fmt.Sprintf("Failed to insert first place score - %v", err))
				handlers.Return500(c)
				return err
			}

			gainedFirstPlace = true
		default:
			h.logError(fmt.Sprintf("Failed to fetch existing first place score - %v", err))
			handlers.Return500(c)
			return nil
		}
	} else {
		ok, err := h.updateFirstPlaceScore(&existingFp)

		if err != nil {
			h.logError(fmt.Sprintf("Failed to update existing first place score -%v", err))
			handlers.Return500(c)
			return err
		}

		if ok {
			gainedFirstPlace = true	
		}
	}

	if !gainedFirstPlace {
		return nil
	}

	mapStr := h.mapData.GetString()

	// Update Activity Feed For First Place Winner
	err = db.InsertActivityFeed(h.user.Id, db.ActivityFeedAchievedFirstPlace, mapStr, h.mapData.MapsetId)

	if err != nil {
		h.logError(fmt.Sprintf("Failed to insert first place winner activity feed - %v", err))
		handlers.Return500(c)
		return err
	}

	// Update activity feed for the first place loser
	if existingFp != (db.FirstPlaceScore{}) && existingFp.UserId != h.user.Id {
		err = db.InsertActivityFeed(existingFp.UserId, db.ActivityFeedLostFirstPlace, mapStr, h.mapData.MapsetId)

		if err != nil {
			h.logError(fmt.Sprintf("Failed to insert first place loser activity feed - %v", err))
			handlers.Return500(c)
			return err
		}
	}

	var oldUser db.User
	oldUser, err = db.GetUserById(existingFp.UserId)

	if err != nil && err != sql.ErrNoRows {
		h.logError(fmt.Sprintf("Failed to fetch existing first place - %v", err))
		handlers.Return500(c)
		return err
	}

	dbScore := h.convertToDbScore()
	err = utils.SendFirstPlaceWebhook(&h.user, &dbScore, &h.mapData, &oldUser)

	// No need to return a 500 here, that could be an issue on Discord's end, it's not crucial to log webhooks.
	if err != nil {
		h.logError(fmt.Sprintf("Failed to send first place score to Discord - %v", err))
	}

	return nil
}

// Inserts a new first place score into the database.
func (h *Handler) insertFirstPlaceScore() error {
	fp := db.NewFirstPlaceScore(h.mapData.MD5, h.user.Id, int(h.newScoreId), h.rating.Rating)

	err := fp.Insert()

	if err != nil {
		return nil
	}

	return nil
}

// Updates an existing first place score to the new user. Returns if a user beat the first place score
// or not
func (h *Handler) updateFirstPlaceScore(score *db.FirstPlaceScore) (bool, error){
	if h.rating.Rating < score.PerformanceRating {
		return false, nil
	}

	fp := db.NewFirstPlaceScore(h.mapData.MD5, h.user.Id, int(h.newScoreId), h.rating.Rating)

	err := fp.Update()

	if err != nil {
		return false, err
	}

	return true, err
}

// Checks and unlocks achievements
func (h *Handler) unlockAchievements(c *gin.Context) error {
	score := h.convertToDbScore()
	
	var err error
	h.unlockedAchievements, err = achievements.CheckAchievementsWithNewScore(&h.user, &score, &h.stats)
	
	if err != nil {
		h.logError(fmt.Sprintf("Failed while unlocking achievements - %v", err))
		handlers.Return500(c)
		return err
	}
	
	return nil
}

// Updates elastic search on the API. This is ran in a goroutine because the result
// isn't important enough to block score submission.
func (h *Handler) updateElasticSearch() {
	go func() {
		err := utils.UpdateElasticSearchMapset(h.mapData.MapsetId)

		if err != nil {
			h.logError(fmt.Sprintf("Failed while updating ElasticSearch - %v", err))
		}
	}()
}

// After submitting a score, this will send the user a 200 response
func (h *Handler) sendSuccessfulResponse(c *gin.Context) {
	globalRank, err := h.user.GetGlobalRank(h.mapData.GameMode)
	
	if err != nil {
		h.logError(fmt.Sprintf("Failed to retrieve user global rank - %v", err))
		globalRank = -1
	}
	
	countryRank, err := h.user.GetCountryRank(h.mapData.GameMode)

	if err != nil {
		h.logError(fmt.Sprintf("Failed to retrieve user country rank - %v", err))
		countryRank = -1
	}
	
	status := http.StatusOK
	
	c.JSON(status, gin.H {
		"status": status,
		"timestamp": time.Now().UnixNano() / int64(time.Millisecond),
		"game_mode": h.mapData.GameMode,
		"map": gin.H {
			"id": h.mapData.Id,
			"md5": h.mapData.MD5,
		},
		"stats": gin.H {
			"new_global_rank": globalRank,
			"new_country_rank": countryRank,
			"total_score": h.stats.TotalScore,
			"ranked_score": h.stats.RankedScore,
			"overall_accuracy": h.stats.OverallAccuracy,
			"overall_performance_rating": h.stats.OverallRating,
			"play_count": h.stats.PlayCount,
		},
		"score": gin.H {
			"personal_best": h.isPersonalBestScore(),
			"performance_rating": h.rating.Rating,
			"rank": -1,
		},
		"achievements": h.unlockedAchievements,
	})
}

// Logs out the score in a readable way
func (h *Handler) logScore(d time.Duration) {
	log.Info(fmt.Sprintf("Score: #%v | User: %v (#%v) | Map: #%v | PR: %.2f | Acc: %.2f%% | Time: %vs",
		h.newScoreId, h.user.Username, h.user.Id, h.mapData.Id, h.rating.Rating, h.scoreData.Accuracy, d.Seconds()))
}

// Converts the incoming score's to a db score.
func (h *Handler) convertToDbScore() db.Score {
	return db.Score{
		Id:                          int(h.newScoreId),
		UserId:                      h.user.Id,
		MapMD5:                      h.mapData.MD5,
		ReplayMD5:                   h.scoreData.ReplayMD5,
		Mode:                        h.scoreData.GameMode,
		PersonalBest:                h.isPersonalBestScore(),
		PerformanceRating:           h.rating.Rating,
		Mods:                        h.scoreData.Mods,
		Failed:                      h.scoreData.Failed,
		TotalScore:                  h.scoreData.TotalScore,
		Accuracy:                    h.scoreData.Accuracy,
		MaxCombo:                    int(h.scoreData.MaxCombo),
		CountMarv:                   int(h.scoreData.CountMarv),
		CountPerf:                   int(h.scoreData.CountMarv),
		CountGreat:                  int(h.scoreData.CountMarv),
		CountGood:                   int(h.scoreData.CountMarv),
		CountOkay:                   int(h.scoreData.CountMarv),
		CountMiss:                   int(h.scoreData.CountMarv),
		Grade:                       common.GetGradeFromAccuracy(h.scoreData.Accuracy, h.scoreData.Failed),
		ScrollSpeed:                 int(h.scoreData.ScrollSpeed),
		TimePlayStart:               h.scoreData.TimePlayStart,
		TimePlayEnd:                 h.scoreData.TimePlayEnded,
		ExecutingAssembly:           h.scoreData.ExecutingAssemblyMD5,
		EntryAssembly:               h.scoreData.EntryAssemblyMD5,
		QuaverVersion:               h.scoreData.ReplayVersion,
		PauseCount:                  int(h.scoreData.PauseCount),
		PerformanceProcessorVersion: h.rating.Version,
		DifficultyProcessorVersion:  h.difficulty.Result.Version,
		IsDonatorScore:              h.mapData.RankedStatus != common.StatusRanked,
	}
}

func (h *Handler) logIgnoringScore(reason string) {
	log.Warning(fmt.Sprintf("Ignoring score from %v: %v", h.user.ToString(), reason))
}

func (h *Handler) logError(reason string) {
	log.Errorf(fmt.Sprintf("Error submitting score from %v: %v", h.user.ToString(), reason))
}
