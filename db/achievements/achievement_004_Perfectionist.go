package achievements

import (
	"database/sql"
	"github.com/Swan/Nameless/db"
)

type AchievementPerfectionist Achievement

func NewAchievementPerfectionist() AchievementPerfectionist {
	return AchievementPerfectionist{
		Id:           4,
		Name:         "Perfectionist",
		SteamAPIName: "PERFECTIONIST",
	}
}

func (a AchievementPerfectionist) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if !score.Failed && !score.IsDonatorScore &&
		score.CountGreat == 0 && score.CountGood == 0 && score.CountOkay == 0 && score.CountMiss == 0 {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db.Score

	q := "SELECT id FROM scores WHERE failed = 0 AND is_donator_score = 0 " +
		"AND count_great = 0 AND count_good = 0 AND count_okay = 0 AND count_miss = 0 " +
		"AND user_id = ? LIMIT 1"

	err := db.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, err
}
