package db

import (
	"database/sql"
	"github.com/Swan/Nameless/src/common"
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

// Helper function to scan a score's row coming from the database.
func scanScore(score *Score, row *sql.Row) error {
	err := row.Scan(
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
