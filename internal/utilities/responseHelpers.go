package utilities

import (
	"encoding/json"
	"net/http"
)

func RespondWithJson(res http.ResponseWriter, statusCode int, payload interface{}) {
	dat, _ := json.Marshal(payload)
	res.WriteHeader(statusCode)
	res.Write(dat)
}

func RespondWithError(res http.ResponseWriter, statusCode int, message string) {
	payload := struct {
		Error string `json:"error"`
	}{
		message,
	}
	RespondWithJson(res, statusCode, payload)
}
