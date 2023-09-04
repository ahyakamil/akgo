package config

import (
	"akgo/aklog"
	"akgo/akmdc"
	"akgo/exception"
	"bytes"
	"context"
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
		uriStr := "uri=" + r.RequestURI
		for key, values := range headers {
			for _, value := range values {
				headersStr += "\"" + key + "\"" + ":" + "\"" + value + "\"" + ", "
			}
		}

		bodyStr := "body=" + string(body)
		r.Body = ioutil.NopCloser(bytes.NewReader(body))

		customResponseWriter := &CustomResponseWriter{
			ResponseWriter: w,
			ResponseData:   []byte{},
		}
		defer func() {
			if err := recover(); err != nil {
				statusCode := 500
				responseStr := "response=" + string(customResponseWriter.ResponseData)
				statusCodeStr := "statusCode=" + strconv.Itoa(statusCode)
				log := ":::" + methodStr + " :::" + uriStr + " :::" + headersStr + " :::" + bodyStr + " :::" + statusCodeStr + " :::" + responseStr
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
		next.DefaultServeMux.ServeHTTP(customResponseWriter, r.WithContext(ctx))

		statusCode := customResponseWriter.StatusCode
		responseStr := "response=" + string(customResponseWriter.ResponseData)
		statusCodeStr := "statusCode=" + strconv.Itoa(statusCode)
		log := ":::" + methodStr + " :::" + uriStr + " :::" + headersStr + " :::" + bodyStr + " :::" + statusCodeStr + " :::" + responseStr
		if statusCode == 0 || (statusCode >= 200 && statusCode < 300) {
			aklog.Info(log)
		} else if statusCode >= 300 && statusCode < 500 {
			aklog.Warn(log)
		} else {
			aklog.Error(log)
		}
	})
}
