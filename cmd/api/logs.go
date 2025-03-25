package main

import (
	"net/http"
	"strconv"
	"time"

	"auth-service-2.0/internal/repository"
	"github.com/go-chi/chi/v5"
)

func (app *app) createLogHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateLogPayload
	if err := ReadJson(r, &payload); err != nil {
		err := WriteJsonError(w, r, http.StatusBadRequest, err)
		if err != nil {
			return
		}
		return
	}

	if err := Validate.Struct(payload); err != nil {
		err := WriteJsonError(w, r, http.StatusBadRequest, err)
		if err != nil {
			return
		}
		return
	}

	// Set default values for request time if not send
	if payload.RequestTimeStamp.IsZero() {
		payload.RequestTimeStamp = time.Now()
	}

	// todo: remove the Fake user
	userId := int64(1)
	log := &repository.Logs{
		UserID:    userId,
		Tags:      []string{payload.Log},
		CreatedAt: payload.RequestTimeStamp,
	}

	ctx := r.Context()
	if err := app.Store.LoginLogs.Create(ctx, log); err != nil {
		http.Error(w, "Failed to create log: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := app.JsonResponse(w, http.StatusCreated, log); err != nil {
		if writeErr := WriteJsonError(w, r, http.StatusInternalServerError, err); writeErr != nil {
			http.Error(w, writeErr.Error(), http.StatusInternalServerError)
		}
	}
}

func (app *app) getLogsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	IdParam := chi.URLParam(r, "requestId")
	LogId, err := strconv.ParseInt(IdParam, 10, 64)
	if err != nil {
		err := WriteJsonError(w, r, http.StatusInternalServerError, err)
		if err != nil {
			return
		}
	}
 
	logs, err := app.Store.LoginLogs.GetById(ctx, int64(LogId))
	if err != nil {
		http.Error(w, "Failed to get logs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := app.JsonResponse(w, http.StatusOK, logs); err != nil {
		if writeErr := WriteJsonError(w, r, http.StatusInternalServerError, err); writeErr != nil {
			http.Error(w, writeErr.Error(), http.StatusInternalServerError)
		}
	}
}
