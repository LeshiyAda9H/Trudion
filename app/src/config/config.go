package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

var DbConfig DBConfig

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	DbConfig = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}
}
