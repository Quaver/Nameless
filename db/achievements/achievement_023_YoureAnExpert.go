package achievements

import (
	"database/sql"
	db2 "github.com/Swan/Nameless/db"
)

type AchievementYoureAnExpert Achievement

func NewAchievementYoureAnExpert() AchievementYoureAnExpert {
	return AchievementYoureAnExpert{
		Id: 23,
		Name: "You're An Expert",
		SteamAPIName: "YOURE_AN_EXPERT",
	}
}

func (a AchievementYoureAnExpert) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	if score.PerformanceRating >= 30 && !score.IsDonatorScore {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db2.Score
	q := "SELECT id FROM scores WHERE performance_rating >= 30 AND is_donator_score = 0 AND user_id = ? LIMIT 1"
	err := db2.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
