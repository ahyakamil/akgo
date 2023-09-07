package auth

import (
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
