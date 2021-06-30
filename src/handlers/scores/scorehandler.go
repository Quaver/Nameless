package scores

import (
	"database/sql"
	"fmt"
	"github.com/Swan/Nameless/src/auth"
	"github.com/Swan/Nameless/src/common"
	"github.com/Swan/Nameless/src/db"
	"github.com/Swan/Nameless/src/handlers"
	"github.com/Swan/Nameless/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	scoreData scoreSubmissionData
	user      db.User
	mapData   db.Map
	mapPath   string
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
func (h Handler) handleSubmission(c *gin.Context) error {
	err := h.checkZeroTotalScore(c)

	if err != nil {
		return err
	}

	err = h.checkDuplicateScore(c)

	if err != nil {
		return err
	}

	return nil
}

// Checks if the score has zero total score (no notes hit whatsoever). These scores
// are ignored because they are considered useless.
func (h Handler) checkZeroTotalScore(c *gin.Context) error {
	if !h.scoreData.isValidTotalScore() {
		handlers.Return400(c)
		return fmt.Errorf("ignoring submitted score with 0 total score")
	}

	return nil
}

// Players can sometimes submit duplicate scores unexpectedly (ex. server restarts, timeouts, etc)
// This checks if the score is a duplicate, and will return a 400 if it is.
func (h Handler) checkDuplicateScore(c *gin.Context) error {
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
