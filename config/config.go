package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT       string
	SERVER_URL string
	DATABASE   string
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("cant read env", err.Error())
	}

	PORT = os.Getenv("PORT")
	SERVER_URL = os.Getenv("SERVER_URL")
	DATABASE = os.Getenv("DATABASE")
}
