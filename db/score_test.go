package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Swan/Nameless/common"
	"github.com/Swan/Nameless/config"
)

func TestInitializeScore(t *testing.T) {
	config.InitializeConfig("../")
	InitializeSQL()
}

// Tests fetching a score by its replay md5 hash and user
func TestGetScoreByReplayMD5(t *testing.T) {
	user, err := GetUserById(1)

	if err != nil {
		t.Fatalf(err.Error())
	}

	const expectedId int = 51
	const expectedMD5 string = "06fda1596f47f3e724aee396390031c4"

	score, err := GetScoreByReplayMD5(&user, expectedMD5)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if score.Id != expectedId || score.ReplayMD5 != expectedMD5 {
		t.Fatalf("Expected score id %v", expectedId)
	}
}

// Tests fetching a user's personal best score on a map
func TestGetPersonalBestScore(t *testing.T) {
	user, err := GetUserById(1)

	if err != nil {
		t.Fatalf(err.Error())
	}

	dbMap, err := GetMapById(2)

	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = GetPersonalBestScore(&user, &dbMap)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

// Tests fetching a user's personal best score, but they do not have one
func TestGetNoPersonalBestScore(t *testing.T) {
	user, err := GetUserById(1)

	if err != nil {
		t.Fatalf(err.Error())
	}

	dbMap, err := GetMapById(1500)

	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = GetPersonalBestScore(&user, &dbMap)

	if err != nil && err != sql.ErrNoRows {
		t.Fatalf(err.Error())
	}
}

func TestGetUserTopScores(t *testing.T) {
	_, err := GetUserTopScores(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCalculateOverallRating(t *testing.T) {
	scores, err := GetUserTopScores(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	_ = CalculateOverallRating(scores)
}

func TestCalculateOverallAccuracy(t *testing.T) {
	scores, err := GetUserTopScores(1, common.ModeKeys4)

	if err != nil {
		t.Fatalf(err.Error())
	}

	_ = CalculateOverallAccuracy(scores)
}

func TestGetClanScores(t *testing.T) {
	TestInitializeScore(t)

	scores, err := GetClanPlayerScores(1, "a821c04f97744bb7e3da3c3ff382988d")

	if err != nil {
		t.Fatal(err)
		return
	}

	for _, s := range scores {
		fmt.Printf("%v - %v\n", s.UserId, s.PerformanceRating)
	}

	TestCloseScore(t)
}
func TestCloseScore(t *testing.T) {
	CloseSQLConnection()
}
