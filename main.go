package main

import (
	"akgo/aklog"
	"akgo/config"
	"akgo/env"
	"akgo/exception"
	"akgo/feature/account"
	"akgo/feature/hello"
	"akgo/response"
	"encoding/json"
	"net/http"
)

func main() {
	customServeMux := &config.CustomServeMux{DefaultServeMux: http.DefaultServeMux}
	mux := config.GlobalMiddleware(customServeMux)

	http.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			response.OkStr("App version "+env.AppVersion, writer)
		} else {
			exception.MethodNotAllowed(writer)
		}
	})

	http.HandleFunc("/accounts", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			var registerReq account.RegisterReq
			decoder := json.NewDecoder(request.Body)
			if err := decoder.Decode(&registerReq); err != nil {
				aklog.Warn(err.Error())
				exception.BadRequest(writer)
				return
			}

			isCreated := account.DoRegister(registerReq)
			if isCreated {
				response.JustOk(writer, http.StatusCreated)
			}
		} else {
			exception.MethodNotAllowed(writer)
		}
	})

	http.HandleFunc("/feeds", func(writer http.ResponseWriter, request *http.Request) {
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
