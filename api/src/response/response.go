package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSONSuccess(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	var response Response
	response.Code = code
	response.Message = message
	response.Data = data
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func JSONError(w http.ResponseWriter, code int, err error, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	var response Response
	response.Code = code
	if err != nil {
		response.Message = err.Error()
	} else {
		response.Message = ""
	}

	response.Data = data
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
