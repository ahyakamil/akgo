package account

import (
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgconn"
)

func DoRegister(req RegisterReq) (pgconn.CommandTag, error) {
	validate := validator.New()
	violation := validate.Struct(req)
	if violation != nil {
		return nil, violation
	}

	account := Account{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	result, err := insert(account)
	return result, err
}
