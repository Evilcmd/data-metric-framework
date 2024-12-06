package apis

import (
	"net/http"

	"github.com/Evilcmd/data-metric-framework/internal/utilities"
)

func Root(res http.ResponseWriter, req *http.Request) {
	utilities.RespondWithJson(res, 200, "Hello World")
}

func CheckHealth(res http.ResponseWriter, req *http.Request) {
	payload := struct {
		Message string `json:"message"`
	}{
		"OK",
	}
	utilities.RespondWithJson(res, 200, payload)
}

func ErrCheck(res http.ResponseWriter, req *http.Request) {
	utilities.RespondWithError(res, http.StatusInternalServerError, "checking error response function")
}
