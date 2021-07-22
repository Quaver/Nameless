package achievements

import (
	"database/sql"
	db2 "github.com/Swan/Nameless/db"
)

type AchievementAbsolutelyMarvelous Achievement

func NewAchievementAbsolutelyMarvelous() AchievementAbsolutelyMarvelous {
	return AchievementAbsolutelyMarvelous{
		Id: 2,
		Name: "Absolutely Marvelous",
		SteamAPIName: "ABSOLUTELY_MARVELOUS",
	}
}

func (a AchievementAbsolutelyMarvelous) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	if !score.Failed && score.Accuracy == 100 && !score.IsDonatorScore {
		return true, nil
	}
	
	if user.CheckedPreviousAchievements {
		return false, nil
	}
	
	var dbScore db2.Score
	q := "SELECT id FROM scores WHERE failed = 0 AND accuracy = 100 AND is_donator_score = 0 AND user_id = ? LIMIT 1"
	err := db2.SQL.QueryRow(q, user.Id).Scan(&dbScore.Id)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		
		return false, err
	}
	
	return true, nil
}
