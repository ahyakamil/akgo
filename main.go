package main

import (
	"akgo/config"
	"akgo/env"
	"akgo/exception"
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
