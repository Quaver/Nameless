package achievements

import (
	db2 "github.com/Swan/Nameless/db"
)

type AchievementHumbleBeginnings Achievement

func NewAchievementHumbleBeginnings() AchievementHumbleBeginnings {
	return AchievementHumbleBeginnings{
		Id: 6,
		Name: "Humble Beginnings",
		SteamAPIName: "HUMBLE_BEGINNINGS",
	}
}

func (a AchievementHumbleBeginnings) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	if stats.OverallRating >= 25 {
		return true, nil
	}

	return false, nil
}