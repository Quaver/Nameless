package db

import (
	"testing"
)

// Tests fetching a score by its replay md5 hash and user
func TestGetScoreByReplayMD5(t *testing.T) {
	InitializeSQL()
	defer CloseSQLConnection()

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
