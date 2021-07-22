package db

import (
	"fmt"
	common2 "github.com/Swan/Nameless/common"
	"github.com/go-redis/redis/v8"
	"strings"
)

// UpdateGlobalLeaderboard Updates the global leaderboard for a user in redis
func UpdateGlobalLeaderboard(u *User, mode common2.Mode, rating float64) error {
	err := Redis.ZAdd(RedisCtx, GetGlobalLeaderboardKey(mode), &redis.Z{
		Score: rating,
		Member: u.Id,
	}).Err()

	if err != nil {
		return err
	}
	
	return nil
}

// UpdateCountryLeaderboard Updates the country leaderboard for a given user and game mode
func UpdateCountryLeaderboard(u *User, mode common2.Mode, rating float64) error {
	err := Redis.ZAdd(RedisCtx, GetCountryLeaderboardKey(u, mode), &redis.Z{
		Score: rating,
		Member: u.Id,
	}).Err()

	if err != nil {
		return err
	}
	
	return nil
}

// GetGlobalLeaderboardKey Gets the global leaderboard redis key for a game mode
func GetGlobalLeaderboardKey(mode common2.Mode) string {
	return fmt.Sprintf("quaver:leaderboard:%v", mode)
}

// GetCountryLeaderboardKey Gets the country leaderboard redis key for a country and mode
func GetCountryLeaderboardKey(u *User, mode common2.Mode) string {
	return fmt.Sprintf("quaver:country_leaderboard:%v:%v", strings.ToLower(u.Country), mode)
}
