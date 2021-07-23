package achievements

import (
	"database/sql"
	"github.com/Swan/Nameless/db"
)

type AchievementGoingInsane Achievement

func NewAchievementGoingInsane() AchievementGoingInsane {
	return AchievementGoingInsane{
		Id:           22,
		Name:         "Going Insane",
		SteamAPIName: "GOING_INSANE",
	}
}

func (a AchievementGoingInsane) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if score.PerformanceRating >= 25 && !score.IsDonatorScore {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db.Score
	q := "SELECT id FROM scores WHERE performance_rating >= 25 AND is_donator_score = 0 AND user_id = ? LIMIT 1"
	err := db.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
