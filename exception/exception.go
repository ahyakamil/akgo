package exception

import (
	"be_news_portal/constant/code"
	"encoding/json"
	"net/http"
)

type BaseErrorData struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func MethodNotAllowed(writer http.ResponseWriter) {
	resp := BaseErrorData{
		Code:    code.METHOD_NOT_ALLOWED,
		Message: "Method not allowed",
	}
	build(resp, writer)
}

func build(resp BaseErrorData, writer http.ResponseWriter) {
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(writer, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusMethodNotAllowed)
	writer.Write(jsonResp)
}
