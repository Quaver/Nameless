package achievements

import (
	"database/sql"
	"github.com/Swan/Nameless/src/db"
)

type AchievementYoureAnExpert Achievement

func NewAchievementYoureAnExpert() AchievementYoureAnExpert {
	return AchievementYoureAnExpert {
		Id: 23,
		SteamAPIName: "YOURE_AN_EXPERT",
	}
}

func (a *AchievementYoureAnExpert) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if score.PerformanceRating >= 30 {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db.Score
	q := "SELECT id FROM scores WHERE performance_rating >= 30 AND user_id = ?"
	err := db.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
