package utils

import (
	"github.com/Swan/Nameless/src/config"
	"github.com/Swan/Nameless/src/db"
	"testing"
)

func TestInitializeElasticSearch(t *testing.T) {
	config.InitializeConfig("../../")
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