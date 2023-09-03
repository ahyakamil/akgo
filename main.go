package main

import (
	"be_news_portal/aklog"
	"be_news_portal/akmdc"
	"be_news_portal/env"
	"be_news_portal/exception"
	"be_news_portal/feature/hello"
	"be_news_portal/response"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strconv"
)

type CustomServeMux struct {
	DefaultServeMux *http.ServeMux
}
type CustomResponseWriter struct {
	http.ResponseWriter
	ResponseData []byte
	StatusCode   int
}

func (c *CustomResponseWriter) Write(data []byte) (int, error) {
	c.ResponseData = append(c.ResponseData, data...)
	return c.ResponseWriter.Write(data)
}

func (c *CustomResponseWriter) WriteHeader(statusCode int) {
	c.StatusCode = statusCode
	c.ResponseWriter.WriteHeader(statusCode)
}

func GlobalMiddleware(next *CustomServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), akmdc.MdcKey, make(akmdc.MDC))
		akmdc.Ctx = ctx
		mdc := akmdc.GetMDC()
		mdc["MDC_GROUP"] = uuid.New().String()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			aklog.Error("Error read body!")
			return
		}

		headers := r.Header
		headersStr := "headers="
		methodStr := "method=" + r.Method
		for key, values := range headers {
			for _, value := range values {
				headersStr += "\"" + key + "\"" + ":" + "\"" + value + "\"" + ", "
			}
		}

		bodyStr := "body=" + string(body)

		customResponseWriter := &CustomResponseWriter{
			ResponseWriter: w,
			ResponseData:   []byte{},
		}
		next.DefaultServeMux.ServeHTTP(customResponseWriter, r.WithContext(ctx))

		statusCode := customResponseWriter.StatusCode
		responseStr := "response=" + string(customResponseWriter.ResponseData)
		statusCodeStr := "statusCode=" + strconv.Itoa(statusCode)
		log := ":::" + methodStr + " :::" + headersStr + " :::" + bodyStr + " :::" + statusCodeStr + " :::" + responseStr
		if statusCode == 0 || (statusCode >= 200 && statusCode < 300) {
			aklog.Info(log)
		} else if statusCode >= 300 && statusCode < 500 {
			aklog.Warn(log)
		} else {
			aklog.Error(log)
		}
	})
}

func main() {
	customServeMux := &CustomServeMux{DefaultServeMux: http.DefaultServeMux}
	mux := GlobalMiddleware(customServeMux)

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
