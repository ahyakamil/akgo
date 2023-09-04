package account

import (
	"akgo/db"
	"context"
	"github.com/jackc/pgconn"
)

func insert(req RegisterReq) (pgconn.CommandTag, error) {
	return db.DeferAutoClose(func() (pgconn.CommandTag, error) {
		result, err := db.Pg().Exec(
			context.Background(),
			`INSERT INTO account (username, email, password) values ($1, $2, $3)`,
			req.Username, req.Email, req.Password)
		return result, err
	})
}
