package res

import (
	"MinersGame/pkg/errs"
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	if err, ok := data.(error); ok {
		if he, ok := err.(*errs.HTTPError); ok {
			statusCode = he.Code
			data = he.Message
		} else {
			if statusCode == 0 {
				statusCode = http.StatusInternalServerError
			}
			data = err.Error()
		}
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
