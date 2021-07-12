package achievements

import (
	"database/sql"
	"github.com/Swan/Nameless/src/db"
)

type AchievementBabySteps Achievement

func NewAchievementBabySteps() AchievementBabySteps {
	return AchievementBabySteps{
		Id: 1,
		SteamAPIName: "BABY_STEPS",
	}
}

// Check Unlocked by passing any map
func (a *AchievementBabySteps) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if !score.Failed {
		return true, nil
	}
	
	if user.CheckedPreviousAchievements {
		return false, nil
	}
	
	q := "SELECT id FROM scores WHERE failed = 0 AND is_donator_score = 0 LIMIT 1"
	err := db.SQL.QueryRow(q).Err()

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		
		return false, err
	}
	
	return true, err
}