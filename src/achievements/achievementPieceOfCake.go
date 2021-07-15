package achievements

import (
	"database/sql"
	"github.com/Swan/Nameless/src/db"
)

type AchievementPieceOfCake Achievement

func NewAchievementPieceOfCake() AchievementPieceOfCake {
	return AchievementPieceOfCake {
		Id: 24,
		SteamAPIName: "PIECE_OF_CAKE",
	}
}

func (a *AchievementPieceOfCake) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if score.PerformanceRating >= 35 {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db.Score
	q := "SELECT id FROM scores WHERE performance_rating >= 35 AND user_id = ?"
	err := db.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
