package util

import (
	"encoding/json"
	"go-simple-crawler/internal/server/constants"
	"log"
	"net/http"
)

func WriteError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", constants.ApplicationJSON)
	w.WriteHeader(http.StatusInternalServerError)

	enc := json.NewEncoder(w)

	err2 := enc.Encode(map[string]interface{}{
		"status":  "error",
		"message": err.Error(),
	})
	if err2 != nil {
		log.Printf("[WARN] cannot marshal output: %s", err.Error())
	}
}

func WriteData(w http.ResponseWriter, status int, data interface{}) {
	if data != nil {
		w.Header().Set("Content-Type", constants.ApplicationJSON)
	}

	w.WriteHeader(status)

	if data == nil {
		return
	}

	enc := json.NewEncoder(w)

	err := enc.Encode(data)
	if err != nil {
		log.Printf("[WARN] cannot marshal output: %s", err.Error())
	}
}
