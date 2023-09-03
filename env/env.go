package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var AppVersion string
var ServerPort string
var ServerHost string

func init()  {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	AppVersion = os.Getenv("APP_VERSION")
	ServerPort = os.Getenv("SERVER_PORT")
	ServerHost = os.Getenv("SERVER_HOST")
}


