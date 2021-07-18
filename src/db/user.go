package db

import (
	"fmt"
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

func (u *User) ToString() string {
	return fmt.Sprintf("%v (#%v)", u.Username, u.Id)
}


