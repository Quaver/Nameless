package achievements

import (
	db2 "github.com/Swan/Nameless/db"
)

type AchievementKeptYouPlayingHuh Achievement

func NewAchievementKeptYouPlayingHuh() AchievementKeptYouPlayingHuh {
	return AchievementKeptYouPlayingHuh{
		Id: 5,
		Name: "Kept You Playing Huh?",
		SteamAPIName: "KEPT_YOU_PLAYING_HUH",
	}
}

func (a AchievementKeptYouPlayingHuh) Check(user *db2.User, score *db2.Score, stats *db2.UserStats) (bool, error) {
	if stats.PlayCount >= 1000 {
		return true, nil
	}
	
	return false, nil
}