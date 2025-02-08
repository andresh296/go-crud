package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/andresh296/go-crud/tools/utils"
)

type Config struct {
	Database Database  `json:"database"`
	JWT      JWTConfig `json:"jwt"`
}

type Database struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Schema   string `json:"schema"`
}

func Load() Config {
	root, err := utils.FindModuleRoot()
	if err != nil {
		log.Fatal("error read config: ", err)
	}

	path := filepath.Join(root, "/config/default-config.json")
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("error read config: ", err)
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("error unmarshal config: ", err)
	}

	return config
}

type JWTConfig struct {
	SecretKey      string        `json:"jwt_secret_key" : "mi_super_secreto_32_caracteres"`
	ExpirationTime time.Duration `json:"jwt_expiration_time" : "1h"`
}
