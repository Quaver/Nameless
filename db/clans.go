package db

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type ClanScore struct {
	Id              int
	ClanId          int
	Mode            int
	MapMD5          string
	OverallRating   float64
	OverallAccuracy float64
}

// Calculates an individual score on a map from a clan
func CalculateClanMapScore(clan int, md5 string) (ClanScore, error) {
	scores, err := GetClanPlayerScores(clan, md5)

	if err != nil {
		return ClanScore{}, err
	}

	score := ClanScore{
		ClanId:          clan,
		MapMD5:          md5,
		Mode:            int(scores[0].Mode),
		OverallRating:   CalculateOverallRating(scores),
		OverallAccuracy: CalculateOverallAccuracy(scores),
	}

	return score, nil
}

// Calculates the overall rating of a clan
func CalculateClanOverallRating(clan int, mode int) (float64, error) {
	scores, err := GetClanOverallScoresConverted(clan, mode)

	if err != nil {
		return 0, err
	}

	return CalculateOverallRating(scores), nil
}

// Calculates the overall accuracy of a clan
func CalculateClanOverallAccuracy(clan int, mode int) (float64, error) {
	scores, err := GetClanOverallScoresConverted(clan, mode)

	if err != nil {
		return 0, err
	}

	return CalculateOverallAccuracy(scores), nil
}

// Returns a converted db.Score object from []ClanScore
// This is so we can pass it into calculation methods that require that struct
func GetClanOverallScoresConverted(clan int, mode int) ([]Score, error) {
	scores, err := GetClanOverallScores(clan, mode)

	if err != nil {
		return []Score{}, err
	}

	return ClanScoresToScores(scores), nil
}

// Converts an individual ClanScore to db.Score
func (c *ClanScore) ToScore() Score {
	return Score{
		MapMD5:            c.MapMD5,
		PerformanceRating: c.OverallRating,
		Accuracy:          float32(c.OverallAccuracy),
	}
}

// Converts a slice of ClanScore to db.Score
func ClanScoresToScores(clanScores []ClanScore) []Score {
	var scores []Score

	for _, s := range clanScores {
		scores = append(scores, s.ToScore())
	}

	return scores
}

// Retrieves a clan's top scores
func GetClanOverallScores(clan int, mode int) ([]ClanScore, error) {
	query := "SELECT * FROM clan_scores WHERE clan_id = ? AND mode = ? ORDER BY overall_rating DESC"

	rows, err := SQL.Query(query, clan, mode)

	if err != nil {
		return []ClanScore{}, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var scores []ClanScore

	for rows.Next() {
		var score ClanScore

		err := rows.Scan(&score.Id, &score.ClanId, &score.Mode, &score.MapMD5, &score.OverallRating, &score.OverallAccuracy)

		if err != nil {
			return []ClanScore{}, err
		}

		scores = append(scores, score)
	}

	return scores, nil
}

// Inserts a new clan score to the database (or updates it if it already exists)
func InsertClanScore(score *ClanScore) error {
	delQuery := "DELETE FROM clan_scores WHERE clan_id = ? AND map_md5 = ?"

	_, err := SQL.Exec(delQuery, score.ClanId, score.MapMD5)

	if err != nil {
		return err
	}

	insertQuery := "INSERT INTO clan_scores (clan_id, mode, map_md5, overall_rating, overall_accuracy) VALUES (?, ?, ?, ?, ?)"

	_, err = SQL.Exec(insertQuery, score.ClanId, score.Mode, score.MapMD5, score.OverallRating, score.OverallAccuracy)

	if err != nil {
		return err
	}

	return nil
}

// Updates a clan's stats in the database
func UpdateClanStats(clan int, mode int, rating float64, acc float64) error {
	query := "UPDATE clan_stats SET overall_performance_rating = ?, overall_accuracy = ? WHERE clan_id = ? AND mode = ?"

	_, err := SQL.Exec(query, rating, acc, clan, mode)

	if err != nil {
		return err
	}

	return nil
}

// Update a clan's rating in redis
func UpdateClanLeaderboards(clan int, mode int, rating float64) error {
	err := Redis.ZAdd(RedisCtx, fmt.Sprintf("quaver:clan_leaderboard:%v", mode), &redis.Z{
		Score:  rating,
		Member: clan,
	}).Err()

	if err != nil {
		return err
	}

	return nil
}
