package main

import (
	"akgo/aklog"
	"akgo/config"
	"akgo/db"
	"akgo/env"
	"akgo/exception"
	"akgo/feature/auth"
	"akgo/feature/hello"
	"akgo/response"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func main() {
	db.InitPG()
	defer db.Pg.Close()
	if !config.IsRSAPrivateKey(env.PasswordPrivateKey) {
		log.Fatal("Invalid private key!")
	}

	customServeMux := &config.CustomServeMux{DefaultServeMux: http.DefaultServeMux}
	mux := config.GlobalMiddleware(customServeMux)

	http.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			response.OkStr("App version "+env.AppVersion, writer)
		} else {
			exception.MethodNotAllowed(writer)
		}
	})

	http.HandleFunc("/auth/register", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			var registerReq auth.RegisterReq
			decoder := json.NewDecoder(request.Body)
			decoder.DisallowUnknownFields()
			if err := decoder.Decode(&registerReq); err != nil {
				aklog.Warn(err.Error())
				exception.BadRequest(writer)
				return
			}

			_, err := auth.DoRegister(registerReq)
			if err != nil {
				if strings.Contains(err.Error(), "validation") {
					aklog.Warn(err.Error())
					exception.BadRequestWM(err.Error(), writer)
				} else if strings.Contains(err.Error(), "unique constraint") {
					aklog.Warn(err.Error())
					exception.GeneralWarning(err.Error(), writer, http.StatusConflict)
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

	http.HandleFunc("/auth/login", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			var loginReq auth.LoginReq
			decoder := json.NewDecoder(request.Body)
			decoder.DisallowUnknownFields()
			if err := decoder.Decode(&loginReq); err != nil {
				aklog.Warn(err.Error())
				exception.BadRequest(writer)
				return
			}

			loginResp, err := auth.DoLogin(loginReq)
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

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			resp := hello.DoHello(hello.HelloReq{
				Name: "viola",
			})

			jsonResp, _ := json.Marshal(resp)
			response.Ok(jsonResp, writer)
		} else {
			exception.MethodNotAllowed(writer)
		}
	})

	server := http.Server{
		Addr:    env.ServerHost + ":" + env.ServerPort,
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
