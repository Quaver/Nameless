package achievements

import (
	"fmt"
	"github.com/Swan/Nameless/src/db"
	"strconv"
)

type AchievementApproachingTheBlueZenith Achievement

func NewAchievementApproachingTheBlueZenith() AchievementApproachingTheBlueZenith {
	return AchievementApproachingTheBlueZenith {
		Id: 26,
		SteamAPIName: "APPROACHING_THE_BLUE_ZENITH",
	}
}

func (a *AchievementApproachingTheBlueZenith) Check(user *db.User, score *db.Score, stats *db.UserStats) (bool, error) {
	key := fmt.Sprintf("quaver:leaderboard:%v", score.Mode)
	result, err := db.Redis.ZRevRank(db.RedisCtx, key, strconv.Itoa(user.Id)).Result()
	
	if err != nil {
		return false, err
	}
	
	return result < 100, nil
}