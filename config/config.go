package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var PORT string
var SERVER_URL string

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("cant read env", err.Error())
	}

	PORT = os.Getenv("PORT")
	SERVER_URL = os.Getenv("SERVER_URL")
}
