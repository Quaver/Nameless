package achievements

import (
	"database/sql"
	db2 "github.com/Swan/Nameless/db"
)

type AchievementItsGettingHarder Achievement

func NewAchievementItsGettingHarder() AchievementItsGettingHarder {
	return AchievementItsGettingHarder{
		Id: 21,
		Name: "It's Getting Harder",
		SteamAPIName: "ITS_GETTING_HARDER",
	}
}

func (a AchievementItsGettingHarder) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	if score.PerformanceRating >= 15 && !score.IsDonatorScore {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db2.Score
	q := "SELECT id FROM scores WHERE performance_rating >= 15 AND is_donator_score = 0 AND user_id = ? LIMIT 1"
	err := db2.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
