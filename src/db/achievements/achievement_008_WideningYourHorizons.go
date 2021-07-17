package achievements

import "github.com/Swan/Nameless/src/db"

type AchievementWideningYourHorizons Achievement

func NewAchievementWideningYourHorizons() AchievementWideningYourHorizons {
	return AchievementWideningYourHorizons{
		Id: 8,
		SteamAPIName: "WIDENING_YOUR_HORIZONS",
	}
}

func (a *AchievementWideningYourHorizons) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	if stats.OverallRating >= 300 {
		return true, nil
	}

	return false, nil
}