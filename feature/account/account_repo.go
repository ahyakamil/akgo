package account

import (
	"akgo/db"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
)

func insert(req RegisterReq) (pgconn.CommandTag, error) {
	return db.DeferAutoClose(func() (pgconn.CommandTag, error) {
		id, _ := uuid.NewUUID()
		result, err := db.Pg().Exec(
			context.Background(),
			`INSERT INTO account (id, username, email, password) values ($1, $2, $3, $4)`,
			id.String(), req.Username, req.Email, req.Password)
		return result, err
	})
}
