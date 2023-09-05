package account

import (
	"akgo/db"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
)

func insert(account Account) (pgconn.CommandTag, error) {
	id, _ := uuid.NewUUID()
	result, err := db.Pg.Exec(
		context.Background(),
		`INSERT INTO account (id, username, email, password) values ($1, $2, $3, $4)`,
		id.String(), account.Username, account.Email, account.Password)
	return result, err
}
