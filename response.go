package main

import (
	"encoding/json"
	"net/http"
)

func writeResponse(data interface{}, w http.ResponseWriter) {
	result, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
