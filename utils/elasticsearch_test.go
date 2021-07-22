package utils

import (
	config2 "github.com/Swan/Nameless/config"
	db2 "github.com/Swan/Nameless/db"
	"testing"
)

func TestInitializeElasticSearch(t *testing.T) {
	config2.InitializeConfig("../")
	db2.InitializeSQL()
	InitializeDiscordWebhooks()
}

func TestUpdateElasticSearch(t *testing.T) {
	err := UpdateElasticSearchMapset(2)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCloseElasticSearch(t *testing.T) {
	db2.CloseSQLConnection()
}