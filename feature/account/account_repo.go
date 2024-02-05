package account

import (
	"akgo/aklog"
	"akgo/constant/error_message"
	"akgo/db"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
)

func Insert(account Account) (pgconn.CommandTag, string, error) {
	id, _ := uuid.NewUUID()
	result, err := db.Pg.Exec(
		context.Background(),
		`INSERT INTO account (id, name, about, role, mobile, username, email, password) values ($1, $2, $3, $4, $5, $6, $7, $8)`,
		id.String(), account.Name, account.About, account.Role, account.Mobile, account.Username, account.Email, account.Password)
	return result, id.String(), err
}

func GetLogin(account Account) (Account, error) {
	rows, err := db.Pg.Query(context.Background(), "SELECT id, username FROM account WHERE username=$1 AND password=$2", account.Username, account.Password)
	result := Account{}

	count := 0
	for rows.Next() {
		count += 1
		err = rows.Scan(&result.ID, &result.Username)
		if err != nil {
			aklog.Error("Cannot map result getLogin")
		}
	}

	if count == 0 {
		if err != nil {
			aklog.Error(err.Error())
		}
		err = errors.New(error_message.ERROR_DATA_NOT_FOUND)
	}
	return result, err
}
