package achievements

import (
	"github.com/Swan/Nameless/src/db"
)

type AchievementFingerBreaker Achievement

func NewAchievementFingerBreaker() AchievementFingerBreaker {
	return AchievementFingerBreaker{
		Id: 28,
		Name: "Finger Breaker",
		SteamAPIName: "FINGER_BREAKER",
	}
}

func (a AchievementFingerBreaker) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	return stats.RankedScore >= 100000000, nil
}
