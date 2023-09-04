package exception

import (
	"akgo/constant/code"
	"encoding/json"
	"net/http"
)

type BaseErrorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GeneralError(writer http.ResponseWriter) {
	resp := BaseErrorData{
		Code:    code.GENERAL_ERROR,
		Message: "General error!",
	}
	build(resp, writer, http.StatusInternalServerError)
}

func GeneralErrorWM(message string, writer http.ResponseWriter) {
	resp := BaseErrorData{
		Code:    code.GENERAL_ERROR,
		Message: message,
	}
	build(resp, writer, http.StatusInternalServerError)
}

func BadRequest(writer http.ResponseWriter) {
	resp := BaseErrorData{
		Code:    code.GENERAL_WARNING,
		Message: "Bad request!",
	}
	build(resp, writer, http.StatusBadRequest)
}

func MethodNotAllowed(writer http.ResponseWriter) {
	resp := BaseErrorData{
		Code:    code.GENERAL_WARNING,
		Message: "Method not allowed",
	}
	build(resp, writer, http.StatusMethodNotAllowed)
}

func build(resp BaseErrorData, writer http.ResponseWriter, statusCode int) {
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(writer, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(jsonResp)
}
