package auth

import (
	"akgo/aklog"
	"akgo/config"
	"akgo/exception"
	"akgo/feature/account"
	"akgo/helper"
	"akgo/response"
	"encoding/json"
	"net/http"
	"strings"
)

func AuthController() {
	basePath := "/auth"

	http.HandleFunc(basePath+"/register", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			var registerReq RegisterReq
			decoder := json.NewDecoder(request.Body)
			decoder.DisallowUnknownFields()
			if err := decoder.Decode(&registerReq); err != nil {
				aklog.Warn(err.Error())
				exception.BadRequest(writer)
				return
			}

			registerReq.Role = string(account.ROLE_USER)
			_, err := DoRegister(registerReq)
			if err != nil {
				if strings.Contains(err.Error(), "validation") {
					aklog.Warn(err.Error())
					exception.BadRequestWM(err.Error(), writer)
				} else if strings.Contains(err.Error(), "unique constraint") {
					aklog.Warn(err.Error())
					exception.GeneralWarning(err.Error(), writer, http.StatusConflict)
				} else if err.Error() == ERROR_MAP_ROLE {
					aklog.Warn(err.Error())
					exception.BadRequestWM(err.Error(), writer)
				} else {
					aklog.Error(err.Error())
					panic(err.Error())
				}
				return
			}

			response.JustOk(writer, http.StatusCreated)
		} else {
			exception.MethodNotAllowed(writer)
		}
	})

	http.HandleFunc(basePath+"/login", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			var loginReq LoginReq
			decoder := json.NewDecoder(request.Body)
			decoder.DisallowUnknownFields()
			if err := decoder.Decode(&loginReq); err != nil {
				aklog.Warn(err.Error())
				exception.BadRequest(writer)
				return
			}

			loginResp, err := DoLogin(loginReq)
			if err != nil {
				exception.Unauthorized(writer)
				return
			}
			jsonResp, _ := json.Marshal(loginResp)
			response.Ok(jsonResp, writer)
		} else {
			exception.MethodNotAllowed(writer)
		}
	})

	http.HandleFunc("/auth/token", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			tokenReq := &TokenReq{}
			err := helper.MapQueryParamsToStruct(request, tokenReq)
			if err != nil {
				exception.GeneralErrorWM(err.Error(), writer)
				return
			}

			claims, err := config.ParseRefreshToken(tokenReq.RefreshToken)
			if err != nil {
				exception.GeneralErrorWM(err.Error(), writer)
				return
			}

			user := config.User{
				ID:       claims.UserID,
				Username: claims.Username,
				Role:     claims.Role,
			}

			accessToken, err := config.CreateAccessToken(user)
			if err != nil {
				exception.GeneralErrorWM(err.Error(), writer)
				return
			}

			refreshToken, err := config.CreateRefreshToken(user)
			if err != nil {
				exception.GeneralErrorWM(err.Error(), writer)
				return
			}

			resp := LoginResp{
				AccountId:    user.ID,
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			}
			jsonResp, _ := json.Marshal(resp)
			response.Ok(jsonResp, writer)
		} else {
			exception.MethodNotAllowed(writer)
		}
	})
}
