package db

import (
	"github.com/Swan/Nameless/common"
	"github.com/Swan/Nameless/config"
	"testing"
)

func TestStatsInitialize(t *testing.T) {
	config.InitializeConfig("../")
	InitializeSQL()
}

func TestGetStatsKeys4(t *testing.T) {
	_, err := GetUserStats(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetStatsKeys7(t *testing.T) {
	_, err := GetUserStats(1, common.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestUpdateStats(t *testing.T) {
	stats, err := GetUserStats(1, common.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	stats.RankedScore = 1234

	err = stats.UpdateDatabase()

	if err != nil {
		t.Fatalf(err.Error())
	}

	statsCheck, err := GetUserStats(1, common.ModeKeys7)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if statsCheck.RankedScore != stats.RankedScore {
		t.Fatalf("expected %v ranked score", stats.RankedScore)
	}
}

func TestStatsClose(t *testing.T) {
	CloseSQLConnection()
}
