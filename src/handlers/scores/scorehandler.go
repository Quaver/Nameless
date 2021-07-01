package scores

import (
	"database/sql"
	"fmt"
	"github.com/Swan/Nameless/src/auth"
	"github.com/Swan/Nameless/src/common"
	"github.com/Swan/Nameless/src/db"
	"github.com/Swan/Nameless/src/handlers"
	"github.com/Swan/Nameless/src/processors"
	"github.com/Swan/Nameless/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Handler struct {
	scoreData       scoreSubmissionData
	user            db.User
	mapData         db.Map
	mapPath         string
	difficulty      processors.DifficultyProcessor
	rating          processors.RatingProcessor
	oldPersonalBest db.Score
	newScoreId      int64
}

func (h Handler) SubmitPOST(c *gin.Context) {
	var err error

	h.user, err = auth.GetInGameUser(c)

	if err != nil {
		fmt.Printf("Could not authenticate user - %v\n", err.Error())
		handlers.ReturnError(c, http.StatusForbidden, err.Error())
		return
	}

	h.scoreData, err = parseScoreSubmissionData(c)

	if err != nil {
		fmt.Printf("Invalid score data - %v\n", err.Error())
		handlers.Return400(c)
		return
	}

	hasRankedMods := common.IsModComboRanked(h.scoreData.Mods)

	if !hasRankedMods {
		fmt.Printf("Unranked mods - %v\n", h.scoreData.Mods)
		handlers.Return400(c)
		return
	}

	h.mapData, err = db.GetMapByMD5(h.scoreData.MapMD5)

	if err != nil {
		fmt.Printf("Failed to fetch map from db - %v\n", err.Error())
		handlers.Return400(c)
		return
	}

	err = h.scoreData.validateGameMode(&h.mapData)

	if err != nil {
		fmt.Printf("Non-matching game modes: - %v\n", err.Error())
		handlers.Return400(c)
		return
	}

	h.mapPath, err = utils.CacheQuaFile(h.mapData)

	if err != nil {
		fmt.Printf("unable to cache map file - %v\n", err.Error())
		handlers.Return500(c)
		return
	}

	err = h.handleSubmission(c)

	// Responses are given to the player inside of handleSubmission, so it's not needed here
	if err != nil {
		fmt.Printf("unable to submit score - %v\n", err.Error())
		return
	}

	handlers.ReturnMessage(c, http.StatusOK, "OK")
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
	
	err = h.updateLeaderboardCache(c)
	
	if err != nil {
		return err
	}
	
	// Anything below this point requires the map to be ranked
	// since there will be updating leaderboards, handling achievements, etc.
	if h.mapData.RankedStatus != common.StatusRanked {
		return nil
	}
	
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

	const errStr = "error while inserting score - %v\n"

	if err != nil {
		fmt.Printf(errStr, err)
		handlers.Return500(c)
		return err
	}

	h.newScoreId, err = result.LastInsertId()

	if err != nil {
		fmt.Printf(errStr, err)
		handlers.Return500(c)
		return err
	}
	
	_, err = db.Redis.Incr(db.RedisCtx, "quaver:total_scores").Result()
	
	if err != nil {
		fmt.Printf("Failed to increment total scores in redis - %v\n", err)
		handlers.Return500(c)
		return err
	}

	return nil
}

// Updates the user's latest activity in the database
func (h *Handler) updateUserLatestActivity(c *gin.Context) error {
	err := db.UpdateUserLatestActivity(h.user.Id)
	
	if err != nil {
		fmt.Printf("error while updating user latest activity - %v", err)
		handlers.Return500(c)
		return err
	}
	
	return nil
}

/// TODO: Uploads the replay to azure storage 
func (h *Handler) uploadReplay(c *gin.Context) error {
	if !h.isPersonalBestScore() && h.scoreData.GameId != -1 {
		return nil
	}
	
	return nil
}

// Updates the play + fail count of the map
func (h *Handler) updateMapPlayCount(c *gin.Context) error {
	err := db.IncrementMapPlayCount(h.mapData.Id, h.scoreData.Failed)
	
	if err != nil {
		fmt.Printf("error while incrementing map play count - %v", err)
		handlers.Return500(c)
		return err
	}
	
	return nil
}

// Updates the top 50 score leaderboard cache
func (h *Handler) updateLeaderboardCache(c *gin.Context) error {
	if h.scoreData.Failed {
		return nil
	}
	
	return nil
}