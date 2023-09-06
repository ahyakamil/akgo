package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var AppVersion string
var ServerPort string
var ServerHost string
var PGHost string
var PGPort string
var PGUsername string
var PGPassword string
var PGDatabase string
var PGMinConn int
var PGMaxConn int
var PGMaxIdleTime int
var PasswordPrivateKey string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	AppVersion = os.Getenv("APP_VERSION")
	ServerPort = os.Getenv("SERVER_PORT")
	ServerHost = os.Getenv("SERVER_HOST")

	// postgres
	PGHost = os.Getenv("PG_HOST")
	PGPort = os.Getenv("PG_PORT")
	PGUsername = os.Getenv("PG_USERNAME")
	PGPassword = os.Getenv("PG_PASSWORD")
	PGDatabase = os.Getenv("PG_DATABASE")
	PGMinConn, _ = strconv.Atoi(os.Getenv("PG_MIN_CONN"))
	PGMaxConn, _ = strconv.Atoi(os.Getenv("PG_MAX_CONN"))
	PGMaxIdleTime, _ = strconv.Atoi(os.Getenv("PG_MAX_IDLE_TIME"))

	PasswordPrivateKey = os.Getenv("PASSWORD_PRIVATE_KEY")
}
