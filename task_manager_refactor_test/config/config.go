package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Dbname    string
	DBurl     string
	SecretKey string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env variables")
	}

	dbname := os.Getenv("DB_NAME")
	dburl := os.Getenv("MONGODB_URI")
	secret := os.Getenv("SECRET_KEY")

	AppConfig = &Config{
		Dbname:    dbname,
		DBurl:     dburl,
		SecretKey: secret,
	}
}
