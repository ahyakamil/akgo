package auth

import (
	"akgo/config"
	"akgo/feature/account"
	"errors"
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
	accountModel := account.Account{
		Name:     req.Name,
		About:    req.About,
		Role:     account.MapStringToRole(req.Role),
		Mobile:   req.Mobile,
		Password: hashPassword,
		Username: req.Username,
		Email:    req.Email,
	}

	if accountModel.Role == account.ROLE_UNKNOWN {
		return nil, errors.New(ERROR_MAP_ROLE)
	}

	commandTag, _, err := account.Insert(accountModel)
	return commandTag, err
}

func DoLogin(req LoginReq) (LoginResp, error) {
	resp := LoginResp{}
	validate := validator.New()
	violation := validate.Struct(req)
	if violation != nil {
		return resp, violation
	}

	hashPassword, err := config.HashPassword(req.Password)
	accountModel := account.Account{
		Username: req.Username,
		Password: hashPassword,
	}
	result, err := account.GetLogin(accountModel)
	if err != nil {
		return resp, err
	}

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
