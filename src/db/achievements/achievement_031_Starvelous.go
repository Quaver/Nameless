package achievements

import (
	"database/sql"
	"github.com/Swan/Nameless/src/db"
)

type AchievementStarvelous Achievement

func NewAchievementStarvelous() AchievementStarvelous {
	return AchievementStarvelous{
		Id: 31,
		SteamAPIName: "STARVELOUS",
	}
}

func (a AchievementStarvelous) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if !score.Failed && !score.IsDonatorScore && score.CountMarv == 0 {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db.Score
	q := "SELECT id FROM scores WHERE user_id = ? AND is_donator_score = 0 AND failed = 0 AND count_marv = 0 LIMIT 1"
	err := db.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
