package account

import (
	"akgo/config"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgconn"
)

func DoRegister(req RegisterReq) (pgconn.CommandTag, error) {
	validate := validator.New()
	violation := validate.Struct(req)
	if violation != nil {
		return nil, violation
	}

	hashPassword, _ := config.HashPassword(req.Password)
	account := Account{
		Username: req.Username,
		Email:    req.Email,
		Password: hashPassword,
	}
	result, err := insert(account)
	return result, err
}
