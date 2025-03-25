package main

import "net/http"

func (app *app) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":          "ok",
		"env":             app.AppAddr.Env,
		"current_version": version,
	}

	if err := app.JsonResponse(w, http.StatusOK, data); err != nil {
		if writeErr := WriteJsonError(w, r, http.StatusInternalServerError, err); writeErr != nil {
			http.Error(w, writeErr.Error(), http.StatusInternalServerError)
		}
	}
}
