package utils

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ESConfig struct {
	Address  string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Config struct {
	ESConfig    ESConfig    `yaml:"es"`
	MysqlConfig MysqlConfig `yaml:"mysql"`
}

var config *Config

func MustInitConfig() {
	currDir, _ := os.Getwd()
	confPath := filepath.Join(currDir, "config.yaml")

	log.Printf("Config path: %s", confPath)

	buf, err := os.ReadFile(confPath)
	if err != nil {
		panic(err)
	}

	config = &Config{}
	err = yaml.Unmarshal(buf, config)

	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return config
}
