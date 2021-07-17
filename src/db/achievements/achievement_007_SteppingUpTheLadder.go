package achievements

import "github.com/Swan/Nameless/src/db"

type AchievementSteppingUpTheLadder Achievement

func NewAchievementSteppingUpTheLadder() AchievementSteppingUpTheLadder {
	return AchievementSteppingUpTheLadder{
		Id: 7,
		SteamAPIName: "STEPPING_UP_THE_LADDER",
	}
}

func (a *AchievementSteppingUpTheLadder) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if stats.OverallRating >= 150 {
		return true, nil
	}

	return false, nil
}