package main

import (
	"github.com/Swan/Nameless/src/db"
	"github.com/Swan/Nameless/src/utils"
)

func main() {
	db.InitializeRedis()
	db.InitializeSQL()
	utils.InitializeAzure()
	server{}.Start(8083)
}
