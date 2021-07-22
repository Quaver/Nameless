package achievements

import (
	"database/sql"
	db2 "github.com/Swan/Nameless/db"
)

type AchievementQuombo Achievement

func NewAchievementQuombo() AchievementQuombo {
	return AchievementQuombo{
		Id: 15,
		Name: "Quombo",
		SteamAPIName: "QUOMBO",
	}
}

func (a AchievementQuombo) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	if score.MaxCombo >= 500 && !score.Failed && !score.IsDonatorScore {
		return true, nil
	}
	
	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db2.Score
	q := "SELECT id FROM scores WHERE max_combo >= 500 AND is_donator_score = 0 AND failed = 0 AND user_id = ? LIMIT 1"
	err := db2.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
