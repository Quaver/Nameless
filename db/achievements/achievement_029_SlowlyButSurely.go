package achievements

import (
	"database/sql"
	"fmt"
	"github.com/Swan/Nameless/common"
	"github.com/Swan/Nameless/db"
)

type AchievementSlowlyButSurely Achievement

func NewAchievementSlowlyButSurely() AchievementSlowlyButSurely {
	return AchievementSlowlyButSurely{
		Id:           29,
		Name:         "Slowly But Surely",
		SteamAPIName: "SLOWLY_BUT_SURELY",
	}
}

func (a AchievementSlowlyButSurely) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if !score.Failed && !score.IsDonatorScore && common.IsModActivated(score.Mods, common.ModSpeed05X) {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db.Score
	q := "SELECT id FROM scores WHERE mods & %v != 0 AND user_id = ? AND is_donator_score = 0 AND failed = 0 LIMIT 1"
	err := db.SQL.QueryRow(fmt.Sprintf(q, common.ModSpeed05X), user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
