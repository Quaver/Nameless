package db

import (
	"fmt"
	"github.com/Swan/Nameless/src/common"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Id                          int
	SteamId                     string
	Username                    string
	Allowed                     bool
	Privileges                  int64
	UserGroups                  int64
	LatestActivity              int64
	Country                     string
	CheckedPreviousAchievements bool
	AvatarURL					string
}

// GetUserById Fetches a user object by their id
func GetUserById(id int) (User, error) {
	query := "SELECT " +
		"id, steam_id, username, allowed, privileges, usergroups, latest_activity," +
		"country, checked_previous_achievements, avatar_url " +
		"FROM users WHERE id = ? LIMIT 1"

	var user User

	err := SQL.QueryRow(query, id).Scan(
		&user.Id, &user.SteamId, &user.Username, &user.Allowed,
		&user.Privileges, &user.UserGroups, &user.LatestActivity,
		&user.Country, &user.CheckedPreviousAchievements, &user.AvatarURL)
	
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// UpdateUserLatestActivity Updates a user's latest activity to the current time
func UpdateUserLatestActivity(id int) error {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	_, err := SQL.Exec("UPDATE users SET latest_activity = ? WHERE id = ?", timestamp, id)
	
	if err != nil {
		return err
	}
	
	return err
}

// GetGlobalRank Retrieves a user's global rank for a specific game mode
func (u *User) GetGlobalRank(mode common.Mode) (int64, error) {
	key := fmt.Sprintf("quaver:leaderboard:%v", mode)
	result, err := Redis.ZRevRank(RedisCtx, key, strconv.Itoa(u.Id)).Result()

	if err != nil {
		return -1, err
	}

	return result + 1, nil
}

// GetCountryRank Retrieves a user's country rank for a specific game mode
func (u *User) GetCountryRank(mode common.Mode) (int64, error) {
	key := fmt.Sprintf("quaver:country_leaderboard:%v:%v", strings.ToLower(u.Country), mode)
	result, err := Redis.ZRevRank(RedisCtx, key, strconv.Itoa(u.Id)).Result()

	if err != nil {
		return -1, err
	}

	return result + 1, nil
}

func (u *User) ToString() string {
	return fmt.Sprintf("%v (#%v)", u.Username, u.Id)
}


