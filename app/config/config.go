package config

import (
	"encoding/json"
	"os"
	"log"
)

type AppConfig struct {
	Address string `json:"server_address"` // server address
	ContentDir string `json:"content_dir"`// directory of the file system which is accessible to users
	ReadTimeoutSeconds uint `json:"server_read_timeout"`// server read timeout
	WriteTimeoutSeconds uint `json:"server_write_timeout"`// server write timeout
	MaxHeaderBytes int `json:"server_max_header_bytes"`// serve max header bytes
	MaxBodyBytes uint `json:"server_max_body_bytes"`
}

func ReadConfig(file string) AppConfig {
	configFileContent, err := os.ReadFile(file)
	if err != nil {
		panic("Failed to read config from " + file)
	}

	var config AppConfig
	err = json.Unmarshal(configFileContent, &config)

	if err != nil {
		log.Fatalf("[ERR] Config reading error: %s", err.Error())
	}

	return config
}