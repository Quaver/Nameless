package db

import (
	"testing"

	"github.com/Swan/Nameless/config"
)

func TestInitializeClans(t *testing.T) {
	config.InitializeConfig("../")
	InitializeSQL()
}

func TestFetchClanScores(t *testing.T) {
	scores, err := GetClanScores(12, "58e4c90b390730bb1553aa0f3356af24", 10)

	if err != nil {
		t.Fatal(err)
	}

	if len(scores) != 5 {
		t.Fatal("Expected 5 scores")
	}
}

func TestCalculateClanMapScore(t *testing.T) {
	_, err := CalculateClanMapScore(12, "58e4c90b390730bb1553aa0f3356af24")

	if err != nil {
		t.Fatal(err)
	}
}

func TestCloseClans(t *testing.T) {
	CloseSQLConnection()
}
