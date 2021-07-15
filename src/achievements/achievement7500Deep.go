package achievements

import (
	"database/sql"
	"github.com/Swan/Nameless/src/db"
)

type Achievement7500Deep Achievement

func NewAchievement7500Deep() Achievement7500Deep {
	return Achievement7500Deep {
		Id: 18,
		SteamAPIName: "7500_DEEP",
	}
}

func (a *Achievement7500Deep) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if score.MaxCombo >= 7500 {
		return true, nil
	}

	if user.CheckedPreviousAchievements {
		return false, nil
	}

	var dbScore db.Score
	q := "SELECT id FROM scores WHERE max_combo >= 7500 AND user_id = ? LIMIT 1"
	err := db.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
