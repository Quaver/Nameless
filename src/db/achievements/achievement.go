package achievements

import (
	"fmt"
	"github.com/Swan/Nameless/src/db"
)

type Achievement struct {
	Id int `json:"id"`
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
		var a AchievementChecker
		
		switch achievement.Id {
		case 1:
			a = NewAchievementBabySteps()
		case 2:
			a = NewAchievementAbsolutelyMarvelous()
		case 3:
			a = NewAchievementCombolicious()
		case 4:
			a = NewAchievementPerfectionist()
		case 5:
			a = NewAchievementKeptYouPlayingHuh()
		case 6:
			a = NewAchievementHumbleBeginnings()
		case 7:
			a = NewAchievementSteppingUpTheLadder()
		case 8:
			a = NewAchievementWideningYourHorizons()
		case 9:
			a = NewAchievementReachingNewHeights()
		case 10:
			a = NewAchievementOutOfThisWorld()
		case 11:
			a = NewAchievementArea51()
		case 12:
			a = NewAchievementAlien()
		case 13:
			a = NewAchievementAExtraterrestrial()
		case 14:
			a = NewAchievementET()
		case 15:
			a = NewAchievementQuombo()
		case 16:
			a = NewAchievementOneTwoMayweather()
		case 17:
			a = NewAchievementItsOver5000()
		case 18:
			a = NewAchievement7500Deep()
		case 19:
			a = NewAchievementTenThousand()
		case 20:
			a = NewAchievementBeginnersLuck()
		case 21:
			a = NewAchievementItsGettingHarder()
		case 22:
			a = NewAchievementGoingInsane()
		case 23:
			a = NewAchievementYoureAnExpert()
		case 24:
			a = NewAchievementPieceOfCake()
		case 25:
			a = NewAchievementFailureIsAnOption()
		case 26:
			a = NewAchievementApproachingTheBlueZenith()
		case 27:
			a = NewAchievementClickTheArrows()
		case 28:
			a = NewAchievementFingerBreaker()
		case 29:
			a = NewAchievementSlowlyButSurely()
		case 30:
			a = NewAchievementHeWasNumberOne()
		case 31:
			a = NewAchievementStarvelous()
		default:
			return []Achievement{}, fmt.Errorf("achievement %v not implemented", achievement.Id)
		}
		
		ok, err := a.Check(user, score, stats)
		
		if err != nil {
			return []Achievement{}, err
		}
		
		if !ok {
			continue
		}
		
		unlocked = append(unlocked, achievement)
		_, err = db.SQL.Exec("INSERT INTO user_achievements VALUES (?, ?)", user.Id, achievement.Id)
		
		if err != nil {
			return []Achievement{}, err
		}
	}
	
	return unlocked, nil
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