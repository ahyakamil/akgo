package auth

import (
	"akgo/aklog"
	"akgo/db"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
)

func insert(auth Auth) (pgconn.CommandTag, error) {
	id, _ := uuid.NewUUID()
	result, err := db.Pg.Exec(
		context.Background(),
		`INSERT INTO auth (id, username, email, password) values ($1, $2, $3, $4)`,
		id.String(), auth.Username, auth.Email, auth.Password)
	return result, err
}

func getLogin(auth Auth) (Auth, error) {
	rows, err := db.Pg.Query(context.Background(), "SELECT id, username FROM auth WHERE username=$1 AND password=$2", auth.Username, auth.Password)
	result := Auth{}
	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Username)
		if err != nil {
			aklog.Error("Cannot map result getLogin")
		}
	}
	return result, err
}
