package achievements

import (
	db2 "github.com/Swan/Nameless/db"
)

type AchievementArea51 Achievement

func NewAchievementArea51() AchievementArea51 {
	return AchievementArea51{
		Id: 11,
		Name: "Area 51",
		SteamAPIName: "AREA_51",
	}
}

func (a AchievementArea51) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	if stats.OverallRating >= 700 {
		return true, nil
	}

	return false, nil
}