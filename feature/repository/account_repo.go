package repository

import (
	"akgo/aklog"
	"akgo/constant/error_message"
	"akgo/db"
	"akgo/feature/model"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
)

func Insert(accountModel model.Account) (pgconn.CommandTag, string, error) {
	id, _ := uuid.NewUUID()
	result, err := db.Pg.Exec(
		context.Background(),
		`INSERT INTO model (id, name, about, role, mobile, username, email, password) values ($1, $2, $3, $4, $5, $6, $7, $8)`,
		id.String(), accountModel.Name, accountModel.About, accountModel.Role, accountModel.Mobile, accountModel.Username, accountModel.Email, accountModel.Password)
	return result, id.String(), err
}

func GetLogin(accountModel model.Account) (model.Account, error) {
	rows, err := db.Pg.Query(context.Background(), "SELECT id, username FROM model WHERE username=$1 AND password=$2", accountModel.Username, accountModel.Password)
	result := model.Account{}

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
