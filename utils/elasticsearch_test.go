package utils

import (
	config "github.com/Swan/Nameless/config"
	db "github.com/Swan/Nameless/db"
	"testing"
)

func TestInitializeElasticSearch(t *testing.T) {
	config.InitializeConfig("../")
	db.InitializeSQL()
	InitializeDiscordWebhooks()
}

func TestUpdateElasticSearch(t *testing.T) {
	err := UpdateElasticSearchMapset(2)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCloseElasticSearch(t *testing.T) {
	db.CloseSQLConnection()
}