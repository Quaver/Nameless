package achievements

import (
	db "github.com/Swan/Nameless/db"
)

type AchievementHumbleBeginnings Achievement

func NewAchievementHumbleBeginnings() AchievementHumbleBeginnings {
	return AchievementHumbleBeginnings{
		Id: 6,
		Name: "Humble Beginnings",
		SteamAPIName: "HUMBLE_BEGINNINGS",
	}
}

func (a AchievementHumbleBeginnings) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if stats.OverallRating >= 25 {
		return true, nil
	}

	return false, nil
}