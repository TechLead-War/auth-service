package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *app) mount() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer) // Ensure our server doesn't crash and gives a 500 error
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID) // create a unique request id and attach it to the context, for logging and tracking
	router.Use(middleware.RealIP)    // Get the real IP address of the client from request headers

	// Any request that took more than 60 seconds will be timed out,
	// and gives a 504 Gateway Timeout error.
	router.Use(middleware.Timeout(60 * time.Second))

	// Make a group
	router.Route("/v1", func(r chi.Router) {
		r.Get("/ping", app.healthCheckHandler)
	})

	return router
}

func (app *app) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.AppAddr.AppAddr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute, // max time to wait for the next req, if keep-alive connection is enabled
	}

	log.Printf("Server started at %s", app.AppAddr.AppAddr)
	return srv.ListenAndServe()
}
