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

func DoLogin(req LoginReq) (LoginResp, error) {
	resp := LoginResp{}
	validate := validator.New()
	violation := validate.Struct(req)
	if violation != nil {
		return resp, violation
	}

	hashPassword, err := config.HashPassword(req.Password)
	auth := Auth{
		Username: req.Username,
		Password: hashPassword,
	}
	result, err := getLogin(auth)
	accessToken, err := config.CreateAccessToken(config.User{
		ID:       result.ID,
		Username: result.Username,
	})
	refreshToken, err := config.CreateRefreshToken(config.User{
		ID:       result.ID,
		Username: result.Username,
	})
	resp = LoginResp{
		AccountId:    result.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return resp, err
}
