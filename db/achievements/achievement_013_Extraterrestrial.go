package achievements

import (
	"github.com/Swan/Nameless/db"
)

type AchievementExtraterrestrial Achievement

func NewAchievementAExtraterrestrial() AchievementExtraterrestrial {
	return AchievementExtraterrestrial{
		Id:           13,
		Name:         "Extraterrestrial",
		SteamAPIName: "EXTRATERRESTRIAL",
	}
}

func (a AchievementExtraterrestrial) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if stats.OverallRating >= 900 {
		return true, nil
	}

	return false, nil
}
