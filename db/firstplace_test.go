package db

import (
	config "github.com/Swan/Nameless/config"
	"testing"
)

func TestInitializeFirstPlace(t *testing.T) {
	config.InitializeConfig("../")
	InitializeSQL()
}

func TestInsertFirstPlace(t *testing.T) {
	s := FirstPlaceScore{
		MD5: "test-first-place",
		UserId: 1,
		ScoreId: 999999999,
		PerformanceRating: 1,
	}	
	
	err := s.Insert()
	
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCloseFirstPlace(t *testing.T) {
	CloseSQLConnection()
}
