package achievements

import (
	"database/sql"
	db2 "github.com/Swan/Nameless/db"
)

type AchievementStarvelous Achievement

func NewAchievementStarvelous() AchievementStarvelous {
	return AchievementStarvelous{
		Id: 31,
		Name: "Starvelous",
		SteamAPIName: "STARVELOUS",
	}
}

func (a AchievementStarvelous) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	if !score.Failed && !score.IsDonatorScore && score.CountMarv == 0 {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db2.Score
	q := "SELECT id FROM scores WHERE user_id = ? AND is_donator_score = 0 AND failed = 0 AND count_marv = 0 LIMIT 1"
	err := db2.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
