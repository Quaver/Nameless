package db

import (
	"github.com/Swan/Nameless/src/config"
	"testing"
)

func TestInitializeScoreboard(t *testing.T) {
	config.InitializeConfig("../../")
	InitializeSQL()
	InitializeRedis()
}

func TestUpdateScoreboardCache(t *testing.T) {
	user, err := GetUserById(1)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	score, err := GetScoreByReplayMD5(&user, "5f8450b70698083c9f664d59dbc094a9")

	if err != nil {
		t.Fatalf(err.Error())
	}
	
	mapData, err := GetMapByMD5(score.MapMD5)

	if err != nil {
		t.Fatalf(err.Error())
	}
	
	err = UpdateScoreboardCache(&score, &mapData)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCloseScoreboard(t *testing.T) {
	CloseSQLConnection()
}
