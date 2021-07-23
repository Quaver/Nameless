package db

import (
	"github.com/Swan/Nameless/config"
	"testing"
)

func TestInitializeActivityFeed(t *testing.T) {
	config.InitializeConfig("../")
	InitializeSQL()
}

func TestInsertActivityFeed(t *testing.T) {
	err := InsertActivityFeed(1, ActivityFeedDonated, "", -1)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCloseActivityFeed(t *testing.T) {
	CloseSQLConnection()
}
