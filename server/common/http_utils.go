package common

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendHTTPResponse(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(status)

	if status == http.StatusNoContent {
		return
	}

	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("error marshalling payload: %v\n", err)
	}
}
