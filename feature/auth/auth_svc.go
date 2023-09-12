package auth

import (
	"akgo/config"
	"akgo/db"
	"akgo/feature/account"
	"context"
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

	tx, err := db.Pg.Begin(context.Background())
	defer tx.Commit(context.Background())
	hashPassword, _ := config.HashPassword(req.Password)
	authModel := Auth{
		Username: req.Username,
		Email:    req.Email,
		Password: hashPassword,
	}
	insertAuth, authId, err := Insert(authModel, tx)
	if err != nil {
		return nil, err
	}

	accountModel := account.Account{
		Name:   req.Name,
		About:  req.About,
		Role:   account.MapStringToRole(req.Role),
		Mobile: req.Mobile,
		AuthID: authId,
	}

	if accountModel.Role == account.ROLE_UNKNOWN {
		tx.Rollback(context.Background())
		return nil, errors.New(ERROR_MAP_ROLE)
	}

	_, _, err = account.Insert(accountModel, tx)
	if err != nil {
		tx.Rollback(context.Background())
	}
	return insertAuth, err
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
	result, err := GetLogin(auth)
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
