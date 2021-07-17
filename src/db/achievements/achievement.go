package achievements

import "github.com/Swan/Nameless/src/db"

type Achievement struct {
	Id int `json:"id"`
	SteamAPIName string `json:"steam_api_name"`
}

type UserAchievement struct {
	UserId int
	AchievementId int
}

type AchievementChecker interface {
	Check(*db.User, *db.Score, *db.UserStats) (bool, error)
}

// GetUserUnlockedAchievements Retrieves all of the user's unlocked achievements
func GetUserUnlockedAchievements(id int) ([]UserAchievement, error) {
	query := "SELECT * FROM user_achievements WHERE user_id = ?"
	rows, err := db.SQL.Query(query, id)

	if err != nil {
		return []UserAchievement{}, err
	}

	defer rows.Close()

	var achievements []UserAchievement

	for rows.Next() {
		var a UserAchievement
		err = rows.Scan(&a.UserId, &a.AchievementId)

		if err != nil {
			return []UserAchievement{}, err
		}

		err = rows.Err()

		if err != nil {
			return []UserAchievement{}, err
		}

		achievements = append(achievements, a)
	}

	return achievements, nil
}