package achievements

import (
	db2 "github.com/Swan/Nameless/db"
)

type AchievementClickTheArrows Achievement

func NewAchievementClickTheArrows() AchievementClickTheArrows {
	return AchievementClickTheArrows{
		Id: 27,
		Name: "Click. The. Arrows!",
		SteamAPIName: "CLICK_THE_ARROWS",
	}
}

func (a AchievementClickTheArrows) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	return stats.GetTotalHits() >= 1000000, nil
}
