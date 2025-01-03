package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Database Database `json:"database"`
}

type Database struct {
	Driver string `json:"driver"`
	Host string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Schema string `json:"schema"`
}

func Load() Config {
	root,err:=os.Getwd()
	if err != nil{
		log.Fatal("error read config: ",err)
	}

  path := filepath.Join(root,"../config/default-config.json")
	file,err:=os.ReadFile(path)
	if err !=nil{
	log.Fatal("error,falta config",err)
	}
	var config Config// con C es estructura
	err=json.Unmarshal(file,&config)
	if err !=nil{
		log.Fatal("error: unmarshal config",err)
	}
	return config
}
