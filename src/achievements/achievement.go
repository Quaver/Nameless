package achievements

import "github.com/Swan/Nameless/src/db"

type Achievement struct {
	Id int `json:"id"`
	SteamAPIName string `json:"steam_api_name"`
}

type AchievementChecker interface {
	Check(*db.User, *db.Score, *db.UserStats) bool
}