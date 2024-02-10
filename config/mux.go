package config

import (
	"akgo/aklog"
	"akgo/akmdc"
	"akgo/exception"
	"bytes"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "accept, origin, content-type, x-json, x-prototype-version, x-requested-with, authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, UPDATE, PUT, PATCH, DELETE")
		mdc := akmdc.GetMDC()
		mdcUUID := uuid.New().String()
		mdc["MDC_GROUP"] = mdcUUID

		if r.Method == http.MethodOptions {
			w.WriteHeader(204)
			next.DefaultServeMux.ServeHTTP(w, r)
		} else {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				aklog.Error("Error read body!")
				return
			}

			headers := r.Header
			headersStr := "headers="
			methodStr := "method=" + r.Method
			uriStr := "uri=" + r.Host + r.RequestURI
			for key, values := range headers {
				for _, value := range values {
					headersStr += "\"" + key + "\"" + ":" + "\"" + value + "\"" + ", "
				}
			}

			bodyStr := "body=" + string(body)
			r.Body = ioutil.NopCloser(bytes.NewReader(body))

			nonAuthorizedEndpoints := []string{"/auth/register", "/auth/login", "/auth/token", "/info"}
			nonAuthorizedWildcards := []string{"/public/"}
			requestedEndpoint := r.URL.Path
			isNeedAuthorized := true
			for _, endpoint := range nonAuthorizedEndpoints {
				if requestedEndpoint == endpoint {
					isNeedAuthorized = false
					break
				}
			}

			if isNeedAuthorized {
				for _, endpoint := range nonAuthorizedWildcards {
					if strings.Contains(requestedEndpoint, endpoint) {
						isNeedAuthorized = false
						break
					}
				}
			}

			if isNeedAuthorized {
				authorizationHeader := r.Header.Get("Authorization")
				if authorizationHeader == "" {
					exception.Unauthorized(w)
					return
				} else {
					_, err := ParseAccessToken(authorizationHeader)
					if err != nil {
						exception.Unauthorized(w)
						return
					}
				}
			}

			customResponseWriter := &CustomResponseWriter{
				ResponseWriter: w,
				ResponseData:   []byte{},
			}
			defer func() {
				if err := recover(); err != nil {
					statusCode := 500
					responseStr := "response=" + string(customResponseWriter.ResponseData)
					statusCodeStr := "statusCode=" + strconv.Itoa(statusCode)
					log := ":::" + methodStr + " :::" + statusCodeStr + " :::" + uriStr + " :::" + headersStr + " :::" + bodyStr + " :::" + responseStr
					aklog.Error(log)
					if errStr, ok := err.(string); ok {
						if errStr == "" {
							exception.GeneralError(customResponseWriter)
						} else {
							exception.GeneralErrorWM(errStr, customResponseWriter)
						}
					} else {
						exception.GeneralError(customResponseWriter)
					}
				}
			}()

			customResponseWriter.Header().Set("Mdc-Group", mdcUUID)
			next.DefaultServeMux.ServeHTTP(customResponseWriter, r)

			statusCode := customResponseWriter.StatusCode
			responseStr := "response=" + string(customResponseWriter.ResponseData)
			statusCodeStr := "statusCode=" + strconv.Itoa(statusCode)
			log := ":::" + methodStr + " :::" + statusCodeStr + " :::" + uriStr + " :::" + headersStr + " :::" + bodyStr + " :::" + responseStr
			if statusCode == 0 || (statusCode >= 200 && statusCode < 300) {
				aklog.Info(log)
			} else if statusCode >= 300 && statusCode < 500 {
				aklog.Warn(log)
			} else {
				aklog.Error(log)
			}
		}
	})
}
