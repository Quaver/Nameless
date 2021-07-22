package achievements

import (
	"database/sql"
	db2 "github.com/Swan/Nameless/db"
)

type AchievementBeginnersLuck Achievement

func NewAchievementBeginnersLuck() AchievementBeginnersLuck {
	return AchievementBeginnersLuck{
		Id: 20,
		Name: "Beginner's Luck",
		SteamAPIName: "BEGINNERS_LUCK",
	}
}

func (a AchievementBeginnersLuck) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	if score.PerformanceRating >= 10 && !score.IsDonatorScore {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db2.Score
	q := "SELECT id FROM scores WHERE performance_rating >= 10 AND is_donator_score = 0 AND user_id = ? LIMIT 1"
	err := db2.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
