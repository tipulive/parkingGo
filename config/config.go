package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
	// 1. Load .env
	_ = godotenv.Load()

	// 2. Tell Viper to use env vars
	viper.AutomaticEnv()
	// 3. Read YAML config
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	// 4. Expand env vars inside YAML
	for _, key := range viper.AllKeys() {
		val := viper.GetString(key)
		viper.Set(key, os.ExpandEnv(val))
	}
	var config MySQLConfig
	err = viper.UnmarshalKey("mysql", &config)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	return config
}
