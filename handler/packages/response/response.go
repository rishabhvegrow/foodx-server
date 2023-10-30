package response

import (
	"net/http"
	"encoding/json"
)

func writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    if data != nil {
        json.NewEncoder(w).Encode(data)
    }
}
