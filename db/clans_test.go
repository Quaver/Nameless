package db

import (
	"testing"

	"github.com/Swan/Nameless/config"
)

func TestInitializeClans(t *testing.T) {
	config.InitializeConfig("../")
	InitializeSQL()
}

func TestFetchClanPlayerScores(t *testing.T) {
	scores, err := GetClanPlayerScores(12, "58e4c90b390730bb1553aa0f3356af24")

	if err != nil {
		t.Fatal(err)
	}

	if len(scores) != 5 {
		t.Fatal("Expected 5 scores")
	}
}

func TestFetchClanOverallScores(t *testing.T) {
	_, err := GetClanOverallScores(1, 1)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCalculateClanMapScore(t *testing.T) {
	_, err := CalculateClanMapScore(12, "58e4c90b390730bb1553aa0f3356af24")

	if err != nil {
		t.Fatal(err)
	}
}

func TestCalculateClanOverallRating(t *testing.T) {
	_, err := CalculateClanOverallRating(1, 1)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCalculateClanOverallAccuracy(t *testing.T) {
	_, err := CalculateClanOverallAccuracy(1, 1)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCloseClans(t *testing.T) {
	CloseSQLConnection()
}
