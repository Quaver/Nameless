package main

import (
	"github.com/Swan/Nameless/src/db"
)

func main() {
	db.InitializeRedis()
	db.InitializeSQL()
	server{}.Start(8083)
}
