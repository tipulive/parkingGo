package config

import (
	"log"

	"github.com/spf13/viper"
)

type MySQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func LoadConfig() MySQLConfig {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config MySQLConfig
	err = viper.UnmarshalKey("mysql", &config)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	return config
}
