package achievements

import (
	"database/sql"
	db2 "github.com/Swan/Nameless/db"
)

type AchievementHeWasNumberOne Achievement

func NewAchievementHeWasNumberOne() AchievementHeWasNumberOne {
	return AchievementHeWasNumberOne{
		Id: 30,
		Name: "He Was Number One",
		SteamAPIName: "HE_WAS_NUMBER_ONE",
	}
}

func (a AchievementHeWasNumberOne) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	var id int
	q := "SELECT 1 FROM scores_first_place WHERE user_id = ? LIMIT 1"
	err := db2.SQL.QueryRow(q, user.Id).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
