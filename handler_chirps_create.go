package main

import (
	"encoding/json"
	"net/http"
)

type Chirp struct {
	Body string `json:"body"`
}

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	var chirp Chirp
	if err := json.NewDecoder(r.Body).Decode(&chirp); err != nil {
		http.Error(w, `{"error": "Something went wrong"}`, http.StatusInternalServerError)
		return
	}

	if len(chirp.Body) > 280 {
		http.Error(w, `{"error": "Chirp is too long"}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"valid": true}`))
}
