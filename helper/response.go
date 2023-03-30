package helper

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type Failed struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Error  string `json:"error"`
}

func SuccessResponse(w http.ResponseWriter, code int, data interface{}) {
	success := Success{
		Code:   code,
		Status: http.StatusText(code),
		Data:   data,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonData, _ := json.Marshal(success)
	w.Write(jsonData)
}

func FailedResponse(w http.ResponseWriter, code int, err string) {
	failed := Failed{
		Code:   code,
		Status: http.StatusText(code),
		Error:  err,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonData, _ := json.Marshal(failed)
	w.Write(jsonData)
}
