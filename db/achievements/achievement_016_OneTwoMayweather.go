package achievements

import (
	"database/sql"
	db "github.com/Swan/Nameless/db"
)

type AchievementOneTwoMayweather Achievement

func NewAchievementOneTwoMayweather() AchievementOneTwoMayweather {
	return AchievementOneTwoMayweather{
		Id: 16,
		Name: "One-Two Mayweather",
		SteamAPIName: "ONE_TWO_MAYWEATHER",
	}
}

func (a AchievementOneTwoMayweather) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if score.MaxCombo >= 2500 && !score.Failed && !score.IsDonatorScore {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db.Score
	q := "SELECT id FROM scores WHERE max_combo >= 2500 AND is_donator_score = 0 AND failed = 0 AND user_id = ? LIMIT 1"
	err := db.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
