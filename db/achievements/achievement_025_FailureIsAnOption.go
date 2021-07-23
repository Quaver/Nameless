package achievements

import (
	"github.com/Swan/Nameless/db"
)

type AchievementFailureIsAnOption Achievement

func NewAchievementFailureIsAnOption() AchievementFailureIsAnOption {
	return AchievementFailureIsAnOption{
		Id:           25,
		Name:         "Failure Is An Option",
		SteamAPIName: "FAILURE_IS_AN_OPTION",
	}
}

func (a AchievementFailureIsAnOption) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if stats.FailCount >= 1000 {
		return true, nil
	}

	return false, nil
}
