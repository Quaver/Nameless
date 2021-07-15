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
	if score.PerformanceRating >= 35 && !score.IsDonatorScore {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db.Score
	q := "SELECT id FROM scores WHERE performance_rating >= 35 AND is_donator_score = 0 AND user_id = ? LIMIT 1"
	err := db.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
