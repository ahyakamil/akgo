package auth

import (
	"akgo/aklog"
	"akgo/constant/error_message"
	"akgo/db"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func Insert(auth Auth, tx pgx.Tx) (pgconn.CommandTag, string, error) {
	id, _ := uuid.NewUUID()
	result, err := tx.Exec(
		context.Background(),
		`INSERT INTO auth (id, username, email, password) values ($1, $2, $3, $4)`,
		id.String(), auth.Username, auth.Email, auth.Password)
	return result, id.String(), err
}

func GetLogin(auth Auth) (Auth, error) {
	rows, err := db.Pg.Query(context.Background(), "SELECT id, username FROM auth WHERE username=$1 AND password=$2", auth.Username, auth.Password)
	result := Auth{}

	count := 0
	for rows.Next() {
		count += 1
		err = rows.Scan(&result.ID, &result.Username)
		if err != nil {
			aklog.Error("Cannot map result getLogin")
		}
	}

	if count == 0 {
		err = errors.New(error_message.REPO_DATA_NOT_FOUND)
	}
	return result, err
}
