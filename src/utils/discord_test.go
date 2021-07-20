package utils

import (
	"github.com/Swan/Nameless/src/common"
	"github.com/Swan/Nameless/src/config"
	"github.com/Swan/Nameless/src/db"
	"testing"
)

func TestInitializeDiscord(t *testing.T) {
	config.InitializeConfig("../../")
	db.InitializeSQL()
	InitializeDiscordWebhooks()
}

func TestSendFirstPlaceWebhook(t *testing.T) {
	user, err := db.GetUserById(1)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	mapData, err := db.GetMapById(2)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	score := db.Score{
		Id: 1234,
		UserId: user.Id,
		MapMD5: "test-md5",
		ReplayMD5: "test-replay-md5",
		Timestamp: 0,
		Mode: common.ModeKeys4,
		PersonalBest: true,
		PerformanceRating: 42.69,
		Accuracy: 73.21,
		Mods: common.ModMirror | common.ModSpeed12X,
		MaxCombo: 2,
	}
	
	oldUser, err := db.GetUserById(2)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	err = SendFirstPlaceWebhook(&user, &score, &mapData, &oldUser)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestSendScoreSubmissionErrorWebhook(t *testing.T) {
	user, err := db.GetUserById(1)

	if err != nil {
		t.Fatalf(err.Error())
	}
	
	err = SendScoreSubmissionErrorWebhook(&user, "Failure getting stats")

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestSendAnticheatWebhook(t *testing.T) {
	user, err := db.GetUserById(1)

	if err != nil {
		t.Fatalf(err.Error())
	}

	mapData, err := db.GetMapById(2)

	if err != nil {
		t.Fatalf(err.Error())
	}

	err = SendAnticheatWebhook(&user, &mapData, 0, true, "Test")

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCloseDiscord(t *testing.T) {
	db.CloseSQLConnection()
}