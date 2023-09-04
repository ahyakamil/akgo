package response

import (
	"akgo/constant/code"
	"encoding/json"
	"net/http"
)

type BaseDataJson struct {
	Code int             `json:"code"`
	Data json.RawMessage `json:"data"`
}

type BaseDataString struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

func JustOk(writer http.ResponseWriter, statusCode int) {
	resp := BaseDataString{
		Code: code.GENERAL_SUCCESS,
		Data: "OK",
	}
	buildStr(resp, writer, statusCode)
}

func Ok(data json.RawMessage, writer http.ResponseWriter) {
	resp := BaseDataJson{
		Code: code.GENERAL_SUCCESS,
		Data: data,
	}
	build(resp, writer, http.StatusOK)
}

func OkStr(data string, writer http.ResponseWriter) {
	resp := BaseDataString{
		Code: code.GENERAL_SUCCESS,
		Data: data,
	}
	buildStr(resp, writer, http.StatusOK)
}

func build(resp BaseDataJson, writer http.ResponseWriter, statusCode int) {
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(writer, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(jsonResp)
}

func buildStr(resp BaseDataString, writer http.ResponseWriter, statusCode int) {
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(writer, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(jsonResp)
}
