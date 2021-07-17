package achievements

import (
	"database/sql"
	"github.com/Swan/Nameless/src/db"
)

type AchievementHeWasNumberOne Achievement

func NewAchievementHeWasNumberOne() AchievementHeWasNumberOne {
	return AchievementHeWasNumberOne {
		Id: 30,
		SteamAPIName: "HE_WAS_NUMBER_ONE",
	}
}

func (a *AchievementHeWasNumberOne) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	var count int
	q := "SELECT COUNT(*) as count FROM scores_first_place WHERE user_id = ? LIMIT 1"
	err := db.SQL.QueryRow(q, user.Id).Scan(&count)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return count > 0, nil
}
