package achievements

import (
	db2 "github.com/Swan/Nameless/db"
)

type AchievementFingerBreaker Achievement

func NewAchievementFingerBreaker() AchievementFingerBreaker {
	return AchievementFingerBreaker{
		Id: 28,
		Name: "Finger Breaker",
		SteamAPIName: "FINGER_BREAKER",
	}
}

func (a AchievementFingerBreaker) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	return stats.RankedScore >= 100000000, nil
}
