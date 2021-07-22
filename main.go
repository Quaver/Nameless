package main

import (
	config2 "github.com/Swan/Nameless/config"
	db2 "github.com/Swan/Nameless/db"
	processors2 "github.com/Swan/Nameless/processors"
	utils2 "github.com/Swan/Nameless/utils"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})
	cwd, _ := os.Getwd()
	config2.InitializeConfig(cwd)
	processors2.CompileQuaverTools()
	db2.InitializeRedis()
	db2.InitializeSQL()
	utils2.InitializeAzure()
	utils2.InitializeDiscordWebhooks()
	server{}.Start(config2.Data.Port)
}
