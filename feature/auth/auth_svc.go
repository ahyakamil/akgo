package auth

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
	auth := Auth{
		Username: req.Username,
		Email:    req.Email,
		Password: hashPassword,
	}
	result, err := insert(auth)
	return result, err
}
