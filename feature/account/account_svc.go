package account

import "github.com/jackc/pgconn"

func DoRegister(req RegisterReq) (pgconn.CommandTag, error) {
	account := Account{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	result, err := insert(account)
	return result, err
}
