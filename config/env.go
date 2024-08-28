package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string

	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		Host:       getEnv("HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DBUSER", "root"),
		DBPassword: getEnv("DBPassword", "docker"),
		// DBAddress:  fmt.Sprintf("%s:%s", getEnv("HOST", "127.0.0.1"), getEnv("DBPORT", "4482")),
		DBAddress: getEnv("DBAddress", "localhost:4482"),
		DBName:    getEnv("DBName", "ecom"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
