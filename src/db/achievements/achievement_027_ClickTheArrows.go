package achievements

import (
	"github.com/Swan/Nameless/src/db"
)

type AchievementClickTheArrows Achievement

func NewAchievementClickTheArrows() AchievementClickTheArrows {
	return AchievementClickTheArrows{
		Id: 27,
		Name: "Click. The. Arrows!",
		SteamAPIName: "CLICK_THE_ARROWS",
	}
}

func (a AchievementClickTheArrows) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	return stats.GetTotalHits() >= 1000000, nil
}
