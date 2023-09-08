package account

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func Insert(account Account, tx pgx.Tx) (pgconn.CommandTag, string, error) {
	id, _ := uuid.NewUUID()
	result, err := tx.Exec(
		context.Background(),
		`INSERT INTO account (id, name, about, role, mobile, auth_id) values ($1, $2, $3, $4, $5, $6)`,
		id.String(), account.Name, account.About, account.Role, account.Mobile, account.AuthID)
	return result, id.String(), err
}
