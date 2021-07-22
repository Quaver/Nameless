package achievements

import (
	db2 "github.com/Swan/Nameless/db"
)

type AchievementReachingNewHeights Achievement

func NewAchievementReachingNewHeights() AchievementReachingNewHeights {
	return AchievementReachingNewHeights{
		Id: 9,
		Name: "Reaching New Heights",
		SteamAPIName: "REACHING_NEW_HEIGHTS",
	}
}

func (a AchievementReachingNewHeights) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	if stats.OverallRating >= 500 {
		return true, nil
	}

	return false, nil
}