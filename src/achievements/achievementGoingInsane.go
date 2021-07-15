package achievements

import (
	"database/sql"
	"github.com/Swan/Nameless/src/db"
)

type AchievementGoingInsane Achievement

func NewAchievementGoingInsane() AchievementGoingInsane {
	return AchievementGoingInsane {
		Id: 22,
		SteamAPIName: "GOING_INSANE",
	}
}

func (a *AchievementGoingInsane) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if score.PerformanceRating >= 25 {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db.Score
	q := "SELECT id FROM scores WHERE performance_rating >= 25 AND user_id = ?"
	err := db.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
