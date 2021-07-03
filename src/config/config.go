package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Configuration struct {
	Port int `json:"port"`
	Redis ConfigurationRedis `json:"redis"` 
	SQL ConfigurationSQL `json:"sql"`
	Azure ConfigurationAzure `json:"azure"`
	QuaverAPIPath string `json:"quaver_api_path"`
	MapCacheDir string `json:"map_cache_dir"`
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

var Data Configuration

// InitializeConfig Reads the configuration file
func InitializeConfig(dir string) {
	if Data != (Configuration{}) {
		return
	}
	
	file, err := os.Open(filepath.Join(dir, "config.json"))
	
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
}
