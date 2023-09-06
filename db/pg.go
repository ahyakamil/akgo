package db

import (
	"akgo/env"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

var Pg *pgxpool.Pool
var (
	dbURL = "postgresql://" + env.PGUsername + ":" + env.PGPassword + "@" + env.PGHost + ":" + env.PGPort + "/" + env.PGDatabase
)

func init() {
	config, err := pgxpool.ParseConfig("")
	config.ConnConfig.Host = env.PGHost
	config.ConnConfig.User = env.PGUsername
	config.ConnConfig.Password = env.PGPassword
	config.ConnConfig.Database = env.PGDatabase
	config.MaxConns = int32(env.PGMaxConn)
	config.MinConns = int32(env.PGMinConn)
	config.MaxConnIdleTime = time.Duration(int64(env.PGMaxIdleTime))

	if err != nil {
		log.Fatalf("Error parsing connection string: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	Pg = pool
	if err != nil {
		log.Fatalf("Unable to connect to the database: " + err.Error())
	}
}
