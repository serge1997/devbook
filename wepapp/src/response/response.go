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

func JSON(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	var statusCode int = response.Code
	if response.Code == 0 {
		statusCode = 501
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func JSONError(w http.ResponseWriter, err error, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	var response Response
	response.Code = code
	if err.Error() == "" {
		response.Message = "Uknow erro, please try later."
	}
	response.Message = err.Error()
	response.Data = data
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}
