package achievements

import (
	"database/sql"
	"github.com/Swan/Nameless/src/db"
)

type AchievementQuombo Achievement

func NewAchievementQuombo() AchievementQuombo {
	return AchievementQuombo{
		Id: 15,
		SteamAPIName: "QUOMBO",
	}
}

func (a AchievementQuombo) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if score.MaxCombo >= 500 && !score.Failed && !score.IsDonatorScore {
		return true, nil
	}
	
	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db.Score
	q := "SELECT id FROM scores WHERE max_combo >= 500 AND is_donator_score = 0 AND failed = 0 AND user_id = ? LIMIT 1"
	err := db.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
