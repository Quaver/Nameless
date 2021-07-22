package utils

import (
	config "github.com/Swan/Nameless/config"
	db "github.com/Swan/Nameless/db"
	"testing"
)

func TestInitialize(t *testing.T) {
	config.InitializeConfig("../")
	db.InitializeSQL()
	InitializeAzure()
}

func TestUploadFile(t *testing.T) {
	err := AzureClient.UploadFile("maps", "random2.qua", []byte("hi"))

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestDownloadMap(t *testing.T) {
	const id int32 = 41104
	m, err := db.GetMapById(id)

	if err != nil {
		t.Fatalf("Could not fetch map `%v` in the databse", id)
		return
	}

	_, err = CacheQuaFile(m)

	if err != nil {
		t.Fatalf(err.Error())
		return
	}
}

func TestDispose(t *testing.T) {
	defer db.CloseSQLConnection()
}
