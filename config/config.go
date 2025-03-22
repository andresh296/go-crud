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
    Auth     AuthConfig `json:"auth"`
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

    var configRaw struct {
        Database Database `json:"database"`
        JWT      struct {
            SecretKey      string `json:"secret_key"`
            ExpirationTime string `json:"expiration_time"`
        } `json:"jwt"`
    }
    
    err = json.Unmarshal(file, &configRaw)
    if err != nil {
        log.Fatal("error unmarshal config: ", err)
    }
    

    duration, err := time.ParseDuration(configRaw.JWT.ExpirationTime)
    if err != nil {
        log.Fatal("error parsing duration: ", err)
    }
    
    config := Config{
        Database: configRaw.Database,
        JWT: JWTConfig{
            SecretKey:      configRaw.JWT.SecretKey,
            ExpirationTime: duration,
        },
    }

    return config
}

type JWTConfig struct {
    SecretKey      string        `json:"secret_key"`
    ExpirationTime time.Duration `json:"expiration_time"`
}

type AuthConfig struct {
    SecretKey string `json:"secret_key"`
    Issuer    string `json:"issuer"`
}