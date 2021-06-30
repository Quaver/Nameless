package db

import "errors"

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
}

// GetUserById Fetches a user object by their id
func GetUserById(id int) (User, error) {
	query := "SELECT " +
		"id, steam_id, username, allowed, privileges, usergroups, latest_activity," +
		"country, checked_previous_achievements " +
		"FROM users WHERE id = ? LIMIT 1"

	var user User

	err := SQL.QueryRow(query, id).Scan(
		&user.Id, &user.SteamId, &user.Username, &user.Allowed,
		&user.Privileges, &user.UserGroups, &user.LatestActivity,
		&user.Country, &user.CheckedPreviousAchievements)

	if err != nil {
		return User{}, errors.New("failed to fetch user from the database")
	}

	return user, nil
}
