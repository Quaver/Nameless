package db

type UserAchievement struct {
	UserId int
	AchievementId int
}

// GetUserUnlockedAchievements Retrieves all of the user's unlocked achievements
func GetUserUnlockedAchievements(id int) ([]UserAchievement, error) {
	query := "SELECT * FROM user_achievements WHERE user_id = ?"
	rows, err := SQL.Query(query, id)
	
	if err != nil {
		return []UserAchievement{}, err
	}
	
	defer rows.Close()
	
	achievements := []UserAchievement{}
	
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