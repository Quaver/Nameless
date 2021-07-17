package achievements

import "github.com/Swan/Nameless/src/db"

type AchievementOutOfThisWorld Achievement

func NewAchievementOutOfThisWorld() AchievementOutOfThisWorld {
	return AchievementOutOfThisWorld{
		Id: 10,
		SteamAPIName: "OUT_OF_THIS_WORLD",
	}
}

func (a *AchievementOutOfThisWorld) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if stats.OverallRating >= 600 {
		return true, nil
	}

	return false, nil
}