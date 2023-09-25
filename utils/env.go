package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv() (string, string, string, string){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed loading .env file")
	}
	DbPassword := os.Getenv("DB_PASSWORD")
	DbUsername := os.Getenv("DB_USERNAME")
	DbProtocol := os.Getenv("DB_PROTOCOL")
	DbAddress := os.Getenv("DB_ADDRESS")
	


	return DbPassword, DbUsername, DbProtocol, DbAddress
}
