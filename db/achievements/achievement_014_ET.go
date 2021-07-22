package achievements

import (
	db2 "github.com/Swan/Nameless/db"
)

type AchievementET Achievement

func NewAchievementET() AchievementET {
	return AchievementET{
		Id: 14,
		Name: "ET",
		SteamAPIName: "ET",
	}
}

func (a AchievementET) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	if stats.OverallRating >= 1000 {
		return true, nil
	}

	return false, nil
}