package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	NatsURL     string
	DatabaseURL string
	JWTSecret   string
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()  // allow environment variables to override

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file: %v", err)
	}

	AppConfig = &Config{
		ServerPort:  viper.GetString("SERVER_PORT"),
		DBHost:      viper.GetString("DB_HOST"),
		DBPort:      viper.GetString("DB_PORT"),
		DBUser:      viper.GetString("DB_USER"),
		DBPassword:  viper.GetString("DB_PASSWORD"),
		DBName:      viper.GetString("DB_NAME"),
		NatsURL:     viper.GetString("NATS_URL"),
		JWTSecret:   viper.GetString("JWT_SECRET"),
	}
}


func get(key, defaultVal string) string {
	if val := viper.GetString(key); val != "" {
		return val
	}
	log.Printf("env %s not set, using default: %s", key, defaultVal)
	return defaultVal
}
