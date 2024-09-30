package controllers

import (
	"encoding/json"
	"net/http"
)

func ExportUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	response := map[string]string{
		"message": "Export process has been initiated.",
		"jobId":   "12345",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
