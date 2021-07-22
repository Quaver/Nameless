package achievements

import (
	db2 "github.com/Swan/Nameless/db"
)

type AchievementExtraterrestrial Achievement

func NewAchievementAExtraterrestrial() AchievementExtraterrestrial {
	return AchievementExtraterrestrial{
		Id: 13,
		Name: "Extraterrestrial",
		SteamAPIName: "EXTRATERRESTRIAL",
	}
}

func (a AchievementExtraterrestrial) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	if stats.OverallRating >= 900 {
		return true, nil
	}

	return false, nil
}