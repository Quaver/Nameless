package db

import (
	"database/sql"
	"github.com/Swan/Nameless/src/common"
)

type Score struct {
	Id int
	UserId int
	MapMD5 string
	ReplayMD5 string
	Timestamp int64
	Mode common.Mode
	PersonalBest bool
	PerformanceRating float64
	Mods common.Mods
	Failed bool
	TotalScore int32
	Accuracy float32
	MaxCombo int
	CountMarv int
	CountPerf int
	CountGreat int
	CountGood int
	CountOkay int
	CountMiss int
	Grade common.Grade
	ScrollSpeed int
	TimePlayStart int64
	TimePlayEnd int64
	Ip string
	ExecutingAssembly string
	EntryAssembly string
	QuaverVersion string
	PauseCount int
	PerformanceProcessorVersion string
	DifficultyProcessorVersion string
	IsDonatorScore bool
	TournamentGameId sql.NullInt32
}

// GetScoreByReplayMD5 Fetches a user's score by replay MD5 in the database
func GetScoreByReplayMD5(u *User, md5 string) (Score, error) {
	query := "SELECT * FROM scores WHERE user_id = ? AND replay_md5 = ?  LIMIT 1"
	
	var score Score
	
	err := SQL.QueryRow(query, u.Id, md5).Scan(
		&score.Id, &score.UserId, &score.MapMD5, &score.ReplayMD5, &score.Timestamp,
		&score.Mode, &score.PersonalBest, &score.PerformanceRating, &score.Mods,
		&score.Failed, &score.TotalScore, &score.Accuracy, &score.MaxCombo,
		&score.CountMarv, &score.CountPerf, &score.CountGreat, &score.CountGood,
		&score.CountOkay, &score.CountMiss, &score.Grade, &score.ScrollSpeed,
		&score.TimePlayStart, &score.TimePlayEnd, &score.Ip, &score.ExecutingAssembly,
		&score.EntryAssembly, &score.QuaverVersion, &score.PauseCount, &score.PerformanceProcessorVersion,
		&score.DifficultyProcessorVersion, &score.IsDonatorScore, &score.TournamentGameId)
	
	if err != nil {
		return Score{}, err
	}
	
	return score, nil
}