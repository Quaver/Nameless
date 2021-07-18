package achievements

import "github.com/Swan/Nameless/src/db"

type AchievementET Achievement

func NewAchievementET() AchievementET {
	return AchievementET{
		Id: 14,
		Name: "ET",
		SteamAPIName: "ET",
	}
}

func (a AchievementET) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if stats.OverallRating >= 1000 {
		return true, nil
	}

	return false, nil
}