package utility

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Server   Server  `json:"server"`
	Data_dir Storage `json:"storage"`
}

type Server struct {
	Port       int `json:"port"`
	Tls_config TLS `json:"tls"`
}

type TLS struct {
	Ca_cert string `json:"ca-cert"`
	Ca_key  string `json:"ca-key"`
	Ca_pem  string `json:"ca-pem"`
}

type Storage struct {
	Data_dir  string `json:"directory"`
	DB_File   string `json:"db_path"`
	Hash_cost string `json:"hashing_cost"` // min 4 , max 31
}

func ParseConfig() *Config {
	currentWd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error:", err)
	}
	configFilePath := filepath.Join(filepath.Dir(currentWd), "config", "config.json")
	configFileData, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	var configStruct Config
	err = json.Unmarshal(configFileData, &configStruct)
	if err != nil {
		log.Fatal(err)
	}
	return &configStruct
}

func CheckDir() {
	currentWd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error:", err)
	}
	dataDir := filepath.Join(filepath.Dir(currentWd), "npm_data")
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		err := os.MkdirAll(dataDir, 0755)
		if err != nil {
			log.Fatal("Unable to create data directory")
		}
	}
}
