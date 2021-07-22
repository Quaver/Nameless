package utils

import (
	common2 "github.com/Swan/Nameless/common"
	config2 "github.com/Swan/Nameless/config"
	db2 "github.com/Swan/Nameless/db"
	"testing"
)

func TestInitializeDiscord(t *testing.T) {
	config2.InitializeConfig("../")
	db2.InitializeSQL()
	InitializeDiscordWebhooks()
}

func TestSendFirstPlaceWebhook(t *testing.T) {
	user, err := db2.GetUserById(1)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	mapData, err := db2.GetMapById(2)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	score := db2.Score{
		Id:                1234,
		UserId:            user.Id,
		MapMD5:            "test-md5",
		ReplayMD5:         "test-replay-md5",
		Timestamp:         0,
		Mode:              common2.ModeKeys4,
		PersonalBest:      true,
		PerformanceRating: 42.69,
		Accuracy:          73.21,
		Mods:              common2.ModMirror | common2.ModSpeed12X,
		MaxCombo:          2,
	}
	
	oldUser, err := db2.GetUserById(2)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	err = SendFirstPlaceWebhook(&user, &score, &mapData, &oldUser)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestSendScoreSubmissionErrorWebhook(t *testing.T) {
	user, err := db2.GetUserById(1)

	if err != nil {
		t.Fatalf(err.Error())
	}
	
	err = SendScoreSubmissionErrorWebhook(&user, "Failure getting stats")

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestSendAnticheatWebhook(t *testing.T) {
	user, err := db2.GetUserById(1)

	if err != nil {
		t.Fatalf(err.Error())
	}

	mapData, err := db2.GetMapById(2)

	if err != nil {
		t.Fatalf(err.Error())
	}

	err = SendAnticheatWebhook(&user, &mapData, 0, true, "Test")

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCloseDiscord(t *testing.T) {
	db2.CloseSQLConnection()
}