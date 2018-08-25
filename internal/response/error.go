package response

import (
	"net/http"
	"encoding/json"
)

type ErrorResponse struct {
	ErrorCode string `json:"code"`
	Message string `json:"message"`
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, errorCode string, msg string) {
	error := ErrorResponse{
		ErrorCode:errorCode,
		Message:msg,
	}

	b, err := json.Marshal(error)

	if err != nil {

	}

	w.WriteHeader(statusCode)
	w.Write(b)
}