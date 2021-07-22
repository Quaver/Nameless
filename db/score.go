package db

import (
	"database/sql"
	common "github.com/Swan/Nameless/common"
	"math"
)

type Score struct {
	Id                          int
	UserId                      int
	MapMD5                      string
	ReplayMD5                   string
	Timestamp                   int64
	Mode                        common.Mode
	PersonalBest                bool
	PerformanceRating           float64
	Mods                        common.Mods
	Failed                      bool
	TotalScore                  int32
	Accuracy                    float32
	MaxCombo                    int
	CountMarv                   int
	CountPerf                   int
	CountGreat                  int
	CountGood                   int
	CountOkay                   int
	CountMiss                   int
	Grade                       common.Grade
	ScrollSpeed                 int
	TimePlayStart               int64
	TimePlayEnd                 int64
	Ip                          string
	ExecutingAssembly           string
	EntryAssembly               string
	QuaverVersion               string
	PauseCount                  int
	PerformanceProcessorVersion string
	DifficultyProcessorVersion  string
	IsDonatorScore              bool
	TournamentGameId            sql.NullInt32
}

// GetScoreByReplayMD5 Fetches a user's score by replay MD5 in the database
func GetScoreByReplayMD5(u *User, md5 string) (Score, error) {
	query := "SELECT * FROM scores WHERE user_id = ? AND replay_md5 = ?  LIMIT 1"

	var score Score

	row := SQL.QueryRow(query, u.Id, md5)
	err := scanScore(&score, row)

	if err != nil {
		return Score{}, err
	}

	return score, nil
}

// GetPersonalBestScore Fetches a user's personal best score on a map
func GetPersonalBestScore(u *User, m *Map) (Score, error) {
	query := "SELECT * FROM scores " +
		"WHERE user_id = ? AND map_md5 = ? AND failed = 0 AND personal_best = 1 " +
		"LIMIT 1"

	var score Score

	row := SQL.QueryRow(query, u.Id, m.MD5)
	err := scanScore(&score, row)

	if err != nil {
		return Score{}, err
	}

	return score, nil
}

// GetUserTopScores Fetches a user's top 500 scores
func GetUserTopScores(id int, mode common.Mode) ([]Score, error) {
	query := "SELECT * FROM scores " +
		"WHERE user_id = ? AND mode = ? AND personal_best = 1 AND is_donator_score = 0 " +
		"ORDER BY performance_rating DESC LIMIT 500"
	
	rows, err := SQL.Query(query, id, mode)
	
	if err != nil {
		return []Score{}, nil
	}
	
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var scores []Score
	
	for rows.Next() {
		var score Score
		err = scanScore(&score, rows)
		
		if err != nil {
			return []Score{}, err
		}
		
		scores = append(scores, score)
	}
	
	return scores, nil
}

// CalculateOverallRating Calculates the overall rating of a list of scores.
// Assumes that scores are sorted by performance rating
func CalculateOverallRating(scores []Score) float64 {
	if len(scores) == 0 {
		return 0
	}
	
	sum := 0.00
	
	for i, score := range scores {
		sum += score.PerformanceRating * math.Pow(0.95, float64(i))
	}
	
	return sum
}

// CalculateOverallAccuracy calculates the overall accuracy from a list of scores
// Assumes that the scores are sorted by performance rating
func CalculateOverallAccuracy(scores []Score) float64 {
	var total float64
	var divideTotal float64
	
	for i, score := range scores {
		add := math.Pow(0.95, float64(i)) * 100
		total += float64(score.Accuracy) * add
		divideTotal += add
	}
	
	if divideTotal == 0 {
		return 0
	}
	
	return total / divideTotal
}

// Helper function to scan a score's row coming from the database.
func scanScore(score *Score, scanner RowScanner) error {
	err := scanner.Scan(
		&score.Id, &score.UserId, &score.MapMD5, &score.ReplayMD5, &score.Timestamp,
		&score.Mode, &score.PersonalBest, &score.PerformanceRating, &score.Mods,
		&score.Failed, &score.TotalScore, &score.Accuracy, &score.MaxCombo,
		&score.CountMarv, &score.CountPerf, &score.CountGreat, &score.CountGood,
		&score.CountOkay, &score.CountMiss, &score.Grade, &score.ScrollSpeed,
		&score.TimePlayStart, &score.TimePlayEnd, &score.Ip, &score.ExecutingAssembly,
		&score.EntryAssembly, &score.QuaverVersion, &score.PauseCount, &score.PerformanceProcessorVersion,
		&score.DifficultyProcessorVersion, &score.IsDonatorScore, &score.TournamentGameId)

	return err
}
