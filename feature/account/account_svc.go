package account

import "github.com/jackc/pgconn"

func DoRegister(req RegisterReq) (pgconn.CommandTag, error) {
	result, err := insert(req)
	return result, err
}
