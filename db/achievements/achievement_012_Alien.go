package achievements

import (
	db2 "github.com/Swan/Nameless/db"
)

type AchievementAlien Achievement

func NewAchievementAlien() AchievementAlien {
	return AchievementAlien{
		Id: 12,
		Name: "Alien",
		SteamAPIName: "ALIEN",
	}
}

func (a AchievementAlien) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	if stats.OverallRating >= 800 {
		return true, nil
	}

	return false, nil
}