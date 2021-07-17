package achievements

import "github.com/Swan/Nameless/src/db"

type Achievement struct {
	Id int `json:"id"`
	SteamAPIName string `json:"steam_api_name"`
}

type AchievementChecker interface {
	Check(*db.User, *db.Score, *db.UserStats) (bool, error)
}

// GetUserUnlockedAchievements Retrieves all of the user's unlocked achievements
func GetUserUnlockedAchievements(id int) ([]Achievement, error) {
	query := "SELECT id, steam_api_name FROM achievements WHERE id IN " +
		"(SELECT achievement_id FROM user_achievements WHERE user_id = ?)"
	
	rows, err := db.SQL.Query(query, id)

	if err != nil {
		return []Achievement{}, err
	}

	defer rows.Close()

	var achievements []Achievement

	for rows.Next() {
		var a Achievement
		err = rows.Scan(&a.Id, &a.SteamAPIName)

		if err != nil {
			return []Achievement{}, err
		}

		err = rows.Err()

		if err != nil {
			return []Achievement{}, err
		}

		achievements = append(achievements, a)
	}

	return achievements, nil
}

// GetUserLockedAchievements Retrieves all of the user's currently locked achievements
func GetUserLockedAchievements(id int) ([]Achievement, error) {
	q :=  "SELECT id, steam_api_name FROM achievements WHERE id NOT IN " +
		"(SELECT achievement_id FROM user_achievements WHERE user_id = ?)"
	
	rows, err := db.SQL.Query(q, id)
	
	if err != nil {
		return []Achievement{}, err
	}
	
	defer rows.Close()
	
	var achievements []Achievement
	
	for rows.Next() {
		var a Achievement
		err = rows.Scan(&a.Id, &a.SteamAPIName)
		
		if err != nil {
			return []Achievement{}, err
		}
		
		err = rows.Err()

		if err != nil {
			return []Achievement{}, err
		}

		achievements = append(achievements, a)
	}
	
	return achievements, nil
}