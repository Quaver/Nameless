package db

import "database/sql"

type ClanScore struct {
	Id              int
	ClanId          int
	MapMD5          string
	OverallRating   float64
	OverallAccuracy float64
}

// Calculates an individual score on a map from a clan
func CalculateClanMapScore(clan int, md5 string) (ClanScore, error) {
	scores, err := GetClanPlayerScores(clan, md5, 10)

	if err != nil {
		return ClanScore{}, err
	}

	score := ClanScore{
		ClanId:          clan,
		OverallRating:   CalculateOverallRating(scores),
		OverallAccuracy: CalculateOverallAccuracy(scores),
	}

	return score, nil
}

// Calculates the overall rating of a clan
func CalculateClanOverallRating(clan int) (float64, error) {
	scores, err := GetClanOverallScoresConverted(clan)

	if err != nil {
		return 0, err
	}

	return CalculateOverallRating(scores), nil
}

// Calculates the overall accuracy of a clan
func CalculateClanOverallAccuracy(clan int) (float64, error) {
	scores, err := GetClanOverallScoresConverted(clan)

	if err != nil {
		return 0, err
	}

	return CalculateOverallAccuracy(scores), nil
}

// Returns a converted db.Score object from []ClanScore
// This is so we can pass it into calculation methods that require that struct
func GetClanOverallScoresConverted(clan int) ([]Score, error) {
	scores, err := GetClanOverallScores(clan)

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
func GetClanOverallScores(clan int) ([]ClanScore, error) {
	query := "SELECT * FROM clan_scores WHERE clan_id = ? ORDER BY overall_rating DESC"

	rows, err := SQL.Query(query, clan)

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

		err := rows.Scan(&score.Id, &score.ClanId, &score.MapMD5, &score.OverallRating, &score.OverallAccuracy)

		if err != nil {
			return []ClanScore{}, err
		}

		scores = append(scores, score)
	}

	return scores, nil
}
