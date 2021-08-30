package db

import (
	"fmt"
	"github.com/Swan/Nameless/common"
)

type UserStats struct {
	Mode              common.Mode
	UserId            int
	TotalScore        int64
	RankedScore       int64
	OverallAccuracy   float64
	OverallRating     float64
	PlayCount         int32
	FailCount         int32
	MaxCombo          int32
	ReplaysWatched    int32
	TotalMarv         int32
	TotalPerf         int32
	TotalGreat        int32
	TotalGood         int32
	TotalOkay         int32
	TotalMiss         int32
	TotalPauses       int
	MultiplayerWins   int
	MultiplayerLosses int
	MultiplayerTies   int
	CountGradeX       int
	CountGradeSS      int
	CountGradeS       int
	CountGradeA       int
	CountGradeB       int
	CountGradeC       int
	CountGradeD       int
}

// GetUserStats Retrieves a user's stats for a game mode
func GetUserStats(id int, mode common.Mode) (UserStats, error) {
	query := fmt.Sprintf("SELECT * FROM %v WHERE user_id = ? LIMIT 1", getTableName(mode))

	stats := UserStats{
		Mode: mode,
	}

	err := SQL.QueryRow(query, id).Scan(
		&stats.UserId, &stats.TotalScore, &stats.RankedScore, &stats.OverallAccuracy, &stats.OverallRating,
		&stats.PlayCount, &stats.FailCount, &stats.MaxCombo, &stats.ReplaysWatched, &stats.TotalMarv,
		&stats.TotalPerf, &stats.TotalGreat, &stats.TotalGood, &stats.TotalOkay, &stats.TotalMiss,
		&stats.TotalPauses, &stats.MultiplayerWins, &stats.MultiplayerLosses, &stats.MultiplayerTies,
		&stats.CountGradeX, &stats.CountGradeSS, &stats.CountGradeS, &stats.CountGradeA, &stats.CountGradeB,
		&stats.CountGradeC, &stats.CountGradeD)

	if err != nil {
		return UserStats{}, err
	}

	return stats, nil
}

// UpdateDatabase Updates the user's stats in the database
func (s *UserStats) UpdateDatabase() error {
	query := fmt.Sprintf("UPDATE %v SET "+
		"total_score = ?, ranked_score = ?, overall_accuracy = ?, overall_performance_rating = ?, "+
		"play_count = ?, fail_count = ?, max_combo = ?, replays_watched = ?, total_marv = ?, "+
		"total_perf = ?, total_great = ?, total_good = ?, total_okay = ?, total_miss = ?, "+
		"total_pauses = ?, multiplayer_wins = ?, multiplayer_losses = ?, multiplayer_ties = ?, " +
		"count_grade_x = ?, count_grade_ss = ?, count_grade_s = ?, count_grade_a = ?, count_grade_b = ?, " +
		"count_grade_c = ?, count_grade_d = ? "+
		"WHERE user_id = ?",
		getTableName(s.Mode))

	_, err := SQL.Exec(query,
		s.TotalScore, s.RankedScore, s.OverallAccuracy, s.OverallRating,
		s.PlayCount, s.FailCount, s.MaxCombo, s.ReplaysWatched, s.TotalMarv,
		s.TotalPerf, s.TotalGreat, s.TotalGood, s.TotalOkay, s.TotalMiss,
		s.TotalPauses, s.MultiplayerWins, s.MultiplayerLosses, s.MultiplayerTies,
		s.CountGradeX, s.CountGradeSS, s.CountGradeS, s.CountGradeA, s.CountGradeB,
		s.CountGradeC, s.CountGradeD,
		s.UserId)

	if err != nil {
		return err
	}

	return nil
}

// UpdateGradeCount Takes a users current pb and old pb and updates the grade count statistic.
func (s *UserStats) UpdateGradeCount(pb Score, oldPb Score) error {
	pbGradeCount, err  := s.getCountGradeFromGrade(common.GetGradeFromAccuracy(pb.Accuracy, pb.Failed))
	
	if err != nil {
		return err
	}
	
	*pbGradeCount++
	
	if oldPb != (Score{}) {
		oldPbGradeCount, err := s.getCountGradeFromGrade(common.GetGradeFromAccuracy(oldPb.Accuracy, oldPb.Failed))
		
		if err != nil {
			return err
		}
		
		*oldPbGradeCount--
	} 
	
	return nil
}

// Returns the "CountGrade" stat from a passed in grade.
func (s *UserStats) getCountGradeFromGrade(grade common.Grade) (*int, error) {
	switch grade {
	case common.GradeX:
		return &s.CountGradeX, nil
	case common.GradeSS:
		return &s.CountGradeSS, nil
	case common.GradeS:
		return &s.CountGradeS, nil
	case common.GradeA:
		return &s.CountGradeA, nil
	case common.GradeB:
		return &s.CountGradeB, nil
	case common.GradeC:
		return &s.CountGradeC, nil
	case common.GradeD:
		return &s.CountGradeD, nil
	default:
		return nil, fmt.Errorf("cannot fetch statistic for grade %v", grade)
	}
}

// GetTotalHits Returns a sum of all hits (misses do onot count)
func (s *UserStats) GetTotalHits() int32 {
	return s.TotalMarv + s.TotalPerf + s.TotalGreat + s.TotalGood + s.TotalOkay
}

// Returns the name of the table depending on the game mode
func getTableName(mode common.Mode) string {
	modeStr, _ := common.GetModeString(mode)
	return fmt.Sprintf("user_stats_%v", modeStr)
}
