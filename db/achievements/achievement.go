package achievements

import (
	"github.com/Swan/Nameless/db"
)

type Achievement struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	SteamAPIName string `json:"steam_api_name"`
}

type AchievementChecker interface {
	Check(*db.User, *db.Score, *db.UserStats) (bool, error)
}

// CheckAchievementsWithNewScore Gets unlocked achievements with an incoming score
func CheckAchievementsWithNewScore(user *db.User, score *db.Score, stats *db.UserStats) ([]Achievement, error) {
	unlocked, err := GetUserUnlockedAchievements(user.Id)

	if err != nil {
		return []Achievement{}, nil
	}

	locked, err := GetUserLockedAchievements(user.Id)

	if err != nil {
		return []Achievement{}, nil
	}

	for _, achievement := range locked {
		a := getAchievementFromId(achievement.Id)
		ok, err := a.Check(user, score, stats)

		if err != nil {
			return []Achievement{}, err
		}

		if !ok {
			continue
		}

		unlocked = append(unlocked, achievement)

		// Give user the achievement
		_, err = db.SQL.Exec("INSERT INTO user_achievements VALUES (?, ?)", user.Id, achievement.Id)

		if err != nil {
			return []Achievement{}, err
		}

		// Display the unlocked achievement in their activity feed
		err = db.InsertActivityFeed(user.Id, db.ActivityFeedUnlockedAchievement, achievement.Name, -1)

		if err != nil {
			return []Achievement{}, err
		}
	}

	// Now that the user's achievements have been checked, we no longer need to do database lookups.
	if !user.CheckedPreviousAchievements {
		_, err = db.SQL.Exec("UPDATE users SET checked_previous_achievements = 1 WHERE id = ?", user.Id)

		if err != nil {
			return []Achievement{}, err
		}
	}

	return unlocked, nil
}

// GetUserUnlockedAchievements Retrieves all of the user's unlocked achievements
func GetUserUnlockedAchievements(id int) ([]Achievement, error) {
	query := "SELECT id, name, steam_api_name FROM achievements WHERE id IN " +
		"(SELECT achievement_id FROM user_achievements WHERE user_id = ?)"

	rows, err := db.SQL.Query(query, id)

	if err != nil {
		return []Achievement{}, err
	}

	defer rows.Close()

	var achievements []Achievement

	for rows.Next() {
		var a Achievement
		err = rows.Scan(&a.Id, &a.Name, &a.SteamAPIName)

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
	q := "SELECT id, name, steam_api_name FROM achievements WHERE id NOT IN " +
		"(SELECT achievement_id FROM user_achievements WHERE user_id = ?)"

	rows, err := db.SQL.Query(q, id)

	if err != nil {
		return []Achievement{}, err
	}

	defer rows.Close()

	var achievements []Achievement

	for rows.Next() {
		var a Achievement
		err = rows.Scan(&a.Id, &a.Name, &a.SteamAPIName)

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

// Returns an achievement checker object given an achievement's id
func getAchievementFromId(id int) AchievementChecker {
	switch id {
	case 1:
		return NewAchievementBabySteps()
	case 2:
		return NewAchievementAbsolutelyMarvelous()
	case 3:
		return NewAchievementCombolicious()
	case 4:
		return NewAchievementPerfectionist()
	case 5:
		return NewAchievementKeptYouPlayingHuh()
	case 6:
		return NewAchievementHumbleBeginnings()
	case 7:
		return NewAchievementSteppingUpTheLadder()
	case 8:
		return NewAchievementWideningYourHorizons()
	case 9:
		return NewAchievementReachingNewHeights()
	case 10:
		return NewAchievementOutOfThisWorld()
	case 11:
		return NewAchievementArea51()
	case 12:
		return NewAchievementAlien()
	case 13:
		return NewAchievementAExtraterrestrial()
	case 14:
		return NewAchievementET()
	case 15:
		return NewAchievementQuombo()
	case 16:
		return NewAchievementOneTwoMayweather()
	case 17:
		return NewAchievementItsOver5000()
	case 18:
		return NewAchievement7500Deep()
	case 19:
		return NewAchievementTenThousand()
	case 20:
		return NewAchievementBeginnersLuck()
	case 21:
		return NewAchievementItsGettingHarder()
	case 22:
		return NewAchievementGoingInsane()
	case 23:
		return NewAchievementYoureAnExpert()
	case 24:
		return NewAchievementPieceOfCake()
	case 25:
		return NewAchievementFailureIsAnOption()
	case 26:
		return NewAchievementApproachingTheBlueZenith()
	case 27:
		return NewAchievementClickTheArrows()
	case 28:
		return NewAchievementFingerBreaker()
	case 29:
		return NewAchievementSlowlyButSurely()
	case 30:
		return NewAchievementHeWasNumberOne()
	case 31:
		return NewAchievementStarvelous()
	}

	return nil
}
