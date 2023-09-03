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

func Ok(data json.RawMessage, writer http.ResponseWriter) {
	resp := BaseDataJson{
		Code: code.GENERAL_SUCCESS,
		Data: data,
	}
	build(resp, writer)
}

func OkStr(data string, writer http.ResponseWriter) {
	resp := BaseDataString{
		Code: code.GENERAL_SUCCESS,
		Data: data,
	}
	buildStr(resp, writer)
}

func build(resp BaseDataJson, writer http.ResponseWriter) {
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(writer, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResp)
}

func buildStr(resp BaseDataString, writer http.ResponseWriter) {
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(writer, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonResp)
}
