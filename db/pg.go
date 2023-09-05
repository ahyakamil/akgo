package db

import (
	"akgo/aklog"
	"akgo/env"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

var PgPool *pgxpool.Pool
var Pg *pgxpool.Conn
var (
	dbURL = "postgresql://" + env.PGUsername + ":" + env.PGPassword + "@" + env.PGHost + ":" + env.PGPort + "/" + env.PGDatabase
)

func DeferAutoClose(f func() (pgconn.CommandTag, error)) (pgconn.CommandTag, error) {
	conn, err := PgPool.Acquire(context.Background())
	Pg = conn
	defer conn.Release()
	if err != nil {
		aklog.Error(err.Error())
		panic(err.Error())
	}
	return f()
}

func init() {
	config, err := pgxpool.ParseConfig("")
	config.ConnConfig.Host = env.PGHost
	config.ConnConfig.User = env.PGUsername
	config.ConnConfig.Password = env.PGPassword
	config.ConnConfig.Database = env.PGDatabase
	config.MaxConns = int32(env.PGMaxConn)
	config.MinConns = int32(env.PGMinConn)

	if err != nil {
		log.Fatalf("Error parsing connection string: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	PgPool = pool
	if err != nil {
		log.Fatalf("Unable to connect to the database: " + err.Error())
	}
}
