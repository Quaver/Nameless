package achievements

import (
	"github.com/Swan/Nameless/db"
)

type AchievementSteppingUpTheLadder Achievement

func NewAchievementSteppingUpTheLadder() AchievementSteppingUpTheLadder {
	return AchievementSteppingUpTheLadder{
		Id:           7,
		Name:         "Stepping Up The Ladder",
		SteamAPIName: "STEPPING_UP_THE_LADDER",
	}
}

func (a AchievementSteppingUpTheLadder) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if stats.OverallRating >= 150 {
		return true, nil
	}

	return false, nil
}
