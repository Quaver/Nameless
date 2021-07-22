package achievements

import (
	"database/sql"
	"fmt"
	common2 "github.com/Swan/Nameless/common"
	db2 "github.com/Swan/Nameless/db"
)

type AchievementSlowlyButSurely Achievement

func NewAchievementSlowlyButSurely() AchievementSlowlyButSurely {
	return AchievementSlowlyButSurely{
		Id: 29,
		Name: "Slowly But Surely",
		SteamAPIName: "SLOWLY_BUT_SURELY",
	}
}

func (a AchievementSlowlyButSurely) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	if !score.Failed && !score.IsDonatorScore && common2.IsModActivated(score.Mods, common2.ModSpeed05X) {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db2.Score
	q := "SELECT id FROM scores WHERE mods & %v != 0 AND user_id = ? AND is_donator_score = 0 AND failed = 0 LIMIT 1"
	err := db2.SQL.QueryRow(fmt.Sprintf(q, common2.ModSpeed05X), user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}
	
	return true, nil
}
