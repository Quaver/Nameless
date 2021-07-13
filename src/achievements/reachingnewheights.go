package achievements

import "github.com/Swan/Nameless/src/db"

type AchievementReachingNewHeights Achievement

func NewAchievementReachingNewHeights() AchievementReachingNewHeights {
	return AchievementReachingNewHeights {
		Id: 9,
		SteamAPIName: "REACHING_NEW_HEIGHTS",
	}
}

func (a *AchievementReachingNewHeights) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if stats.OverallRating >= 500 {
		return true, nil
	}

	return false, nil
}