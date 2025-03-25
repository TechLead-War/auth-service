package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"auth-service-2.0/internal/resources"
	"github.com/go-chi/chi/v5"
)

func (app *app) GetSessionById(w http.ResponseWriter, r *http.Request) {
	sessionIDStr := chi.URLParam(r, "sessionId")

	sessionID, err := strconv.ParseInt(sessionIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	session, err := app.Store.Sessions.FetchSessionByID(r.Context(), sessionID)
	if err != nil {
		if errors.Is(err, resources.ErrLogsNotFound) {
			http.Error(w, "Session not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(session)
	if err != nil {
		return
	}
	log.Printf("%v", err)
}
