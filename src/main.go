package main

import (
	"github.com/Swan/Nameless/src/config"
	"github.com/Swan/Nameless/src/db"
	"github.com/Swan/Nameless/src/processors"
	"github.com/Swan/Nameless/src/utils"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})
	cwd, _ := os.Getwd()
	config.InitializeConfig(cwd)
	processors.CompileQuaverTools()
	db.InitializeRedis()
	db.InitializeSQL()
	utils.InitializeAzure()
	utils.InitializeDiscordWebhooks()
	server{}.Start(config.Data.Port)
}
