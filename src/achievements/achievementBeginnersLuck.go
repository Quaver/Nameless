package achievements

import (
	"database/sql"
	"github.com/Swan/Nameless/src/db"
)

type AchievementBeginnersLuck Achievement

func NewAchievementBeginnersLuck() AchievementBeginnersLuck {
	return AchievementBeginnersLuck {
		Id: 20,
		SteamAPIName: "BEGINNERS_LUCK",
	}
}

func (a *AchievementBeginnersLuck) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if score.PerformanceRating >= 10 {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db.Score
	q := "SELECT id FROM scores WHERE performance_rating >= 10 AND user_id = ?"
	err := db.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
