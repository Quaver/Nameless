package achievements

import (
	db "github.com/Swan/Nameless/db"
)

type AchievementAlien Achievement

func NewAchievementAlien() AchievementAlien {
	return AchievementAlien{
		Id: 12,
		Name: "Alien",
		SteamAPIName: "ALIEN",
	}
}

func (a AchievementAlien) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if stats.OverallRating >= 800 {
		return true, nil
	}

	return false, nil
}