package config

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Configuration struct {
	Port int `json:"port"`
	WebsiteUrl string `json:"website_url"`
	CdnUrl string `json:"cdn_url"`
	QuaverAvatar string `json:"quaver_avatar"`
	Redis ConfigurationRedis `json:"redis"` 
	SQL ConfigurationSQL `json:"sql"`
	Azure ConfigurationAzure `json:"azure"`
	QuaverAPIPath string `json:"quaver_api_path"`
	MapCacheDir string `json:"map_cache_dir"`
	DiscordWebhookFirstPlace ConfigurationDiscordWebhook `json:"discord_webhook_first_place"`
	APIBaseUrl string `json:"api_base_url"`
	APISecretKey string `json:"api_secret_key"`
}

type ConfigurationRedis struct {
	Address string `json:"address"`
	Password string `json:"password"`
	DB int `json:"db"`
}

type ConfigurationSQL struct {
	Host string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type ConfigurationAzure struct {
	AccountName string `json:"account_name"`
	AccountKey string `json:"account_key"`
}

type ConfigurationDiscordWebhook struct {
	Id int64 `json:"id"`
	Secret string `json:"secret"`
}

var Data Configuration

// InitializeConfig Reads the configuration file
func InitializeConfig(dir string) {
	if Data != (Configuration{}) {
		return
	}

	path := filepath.Join(dir, "config.json")
	log.Info(fmt.Sprintf("Loading configuration file at: %v", path))
	
	file, err := os.Open(path)
	
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)
	
	if err != nil {
		panic(err)
	}
	
	err = json.Unmarshal(bytes, &Data)
	
	if err != nil {
		panic(err)
	}
	
	log.Info("Configuration file was successfully read!")
}
