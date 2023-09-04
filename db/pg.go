package db

import (
	"akgo/aklog"
	"akgo/env"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

var (
	dbURL = "postgresql://" + env.PGUsername + ":" + env.PGPassword + "@" + env.PGHost + ":" + env.PGPort + "/" + env.PGDatabase
)

func DeferAutoClose(f func() (pgconn.CommandTag, error)) (pgconn.CommandTag, error) {
	return f()
}

func Pg() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		aklog.Error(err.Error())
	}
	return conn
}
