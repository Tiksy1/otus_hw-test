package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Logger     LoggerConf `json:"logger"`
	HTTPServer HTTPConf   `json:"server"`
	Database   DBConf     `json:"database"`
}

func NewConfig(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("can't open config file: %w", err)
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return Config{}, fmt.Errorf("can't decode config: %w", err)
	}
	return config, nil
}

type LoggerConf struct {
	Level    int8   `json:"level"`
	FilePath string `json:"file_path"` // nolint: tagliatelle
}

type HTTPConf struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type DBConf struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	DBName   string `json:"db_name"` // nolint: tagliatelle
}
