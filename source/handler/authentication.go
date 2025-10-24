package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"npm-registry/authentication"
)

func AuthorizeUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	var authBody authentication.LoginRequest
	json.Unmarshal(body, &authBody)

	authResponse := &authentication.LoginResponse{
		Ok:    "true",
		Token: "1234567890",
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(authResponse)
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}
