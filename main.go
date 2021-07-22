package main

import (
	"github.com/Swan/Nameless/config"
	"github.com/Swan/Nameless/db"
	"github.com/Swan/Nameless/processors"
	"github.com/Swan/Nameless/utils"
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
