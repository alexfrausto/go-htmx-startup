package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port      int32
	DbUser    string
	DbPass    string
	DbHost    string
	DbPort    int32
	DbName    string
	DbSSLMode string
}

var config *Config

func Init() {
	loadConfig()
}

func loadConfig() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 3000
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 5432
	}

	dbSSLMode := os.Getenv("DB_SSLMODE")
	if dbSSLMode == "" {
		dbSSLMode = "disable"
	}

	config = &Config{
		Port:      int32(port),
		DbUser:    os.Getenv("DB_USER"),
		DbPass:    os.Getenv("DB_PASS"),
		DbHost:    os.Getenv("DB_HOST"),
		DbPort:    int32(dbPort),
		DbName:    os.Getenv("DB_NAME"),
		DbSSLMode: dbSSLMode,
	}
}

func GetConfig() *Config {
	return config
}
