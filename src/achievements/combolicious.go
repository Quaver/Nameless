package achievements

import (
	"database/sql"
	"github.com/Swan/Nameless/src/db"
)

type AchievementCombolicious Achievement

func NewAchievementCombolicious() AchievementBabySteps {
	return AchievementBabySteps{
		Id: 3,
		SteamAPIName: "COMBOLICIOUS",
	}
}

func (a *AchievementCombolicious) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if !score.Failed && score.MaxCombo >= 1000 && !score.IsDonatorScore {
		return true, nil
	}
	
	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db.Score
	q := "SELECT id FROM scores WHERE failed = 0 AND max_combo >= 1000 AND is_donator_score = 0 AND user_id = ? LIMIT 1"
	err := db.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}