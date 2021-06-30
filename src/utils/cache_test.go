package utils

import (
	"fmt"
	"github.com/Swan/Nameless/src/db"
	"testing"
)

func TestDownloadMap(t *testing.T) {
	db.InitializeSQL()
	defer db.CloseSQLConnection()
	
	const id int32 = 41093
	m, err := db.GetMapById(id)
	
	if err != nil {
		t.Fatalf("Could not fetch map `%v` in the databse", id)
		return
	}
	
	path, err := CacheQuaFile(m)
	
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	
	fmt.Println(path)
}
