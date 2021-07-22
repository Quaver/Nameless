package achievements

import (
	db "github.com/Swan/Nameless/db"
)

type AchievementReachingNewHeights Achievement

func NewAchievementReachingNewHeights() AchievementReachingNewHeights {
	return AchievementReachingNewHeights{
		Id: 9,
		Name: "Reaching New Heights",
		SteamAPIName: "REACHING_NEW_HEIGHTS",
	}
}

func (a AchievementReachingNewHeights) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if stats.OverallRating >= 500 {
		return true, nil
	}

	return false, nil
}