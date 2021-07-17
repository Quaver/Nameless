package db

import (
	"github.com/Swan/Nameless/src/config"
	"testing"
)

func TestInitializeUserAchievement(t *testing.T) {
	config.InitializeConfig("../../")
	InitializeSQL()
	InitializeRedis()
}

func TestGetUserUnlockedAchievements(t *testing.T) {
	achievements, err := GetUserUnlockedAchievements(1)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	if len(achievements) != 23 {
		t.Fatalf("expected 23 achievement count")
	}
}

func TestCloseUserAchievement(t *testing.T) {
	CloseSQLConnection()
}


