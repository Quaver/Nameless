package achievements

import (
	"database/sql"
	db "github.com/Swan/Nameless/db"
)

type AchievementItsOver5000 Achievement

func NewAchievementItsOver5000() AchievementItsOver5000 {
	return AchievementItsOver5000{
		Id: 17,
		Name: "It's Over 5000!",
		SteamAPIName: "ITS_OVER_5000",
	}
}

func (a AchievementItsOver5000) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if score.MaxCombo >= 5000 && !score.Failed && !score.IsDonatorScore {
		return true, nil
	}
	
	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db.Score
	q := "SELECT id FROM scores WHERE max_combo >= 5000 AND is_donator_score = 0 AND failed = 0 AND user_id = ? LIMIT 1"
	err := db.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
