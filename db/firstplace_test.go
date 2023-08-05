package db

import (
	"github.com/Swan/Nameless/config"
	"testing"
)

func TestInitializeFirstPlace(t *testing.T) {
	config.InitializeConfig("../")
	InitializeSQL()
}

func TestInsertFirstPlace(t *testing.T) {
	s := FirstPlaceScore{
		MD5:               "test-first-place",
		UserId:            1,
		ScoreId:           999999999,
		PerformanceRating: 1,
	}

	err := s.Insert()

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestPublishFirstPlaceScoreRedis(t *testing.T) {
	config.InitializeConfig("../")
	InitializeSQL()
	InitializeRedis()

	err := PublishFirstPlaceScoreRedis("Swan", "Random Artist", "Random Title", "Random Diff")

	if err != nil {
		t.Fatal(err)
	}

	CloseSQLConnection()
}

func TestCloseFirstPlace(t *testing.T) {
	CloseSQLConnection()
}
