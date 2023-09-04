package account

import (
	"akgo/aklog"
)

func DoRegister(req RegisterReq) bool {
	_, err := insert(req)
	if err != nil {
		aklog.Error(err.Error())
		panic(err.Error())
	}
	return true
}
