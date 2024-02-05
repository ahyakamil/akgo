package main

import (
	"akgo/config"
	"akgo/db"
	"akgo/env"
	"akgo/exception"
	"akgo/feature/auth"
	"akgo/response"
	"log"
	"net/http"
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

	auth.AuthController()
	server := http.Server{
		Addr:    env.ServerHost + ":" + env.ServerPort,
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
