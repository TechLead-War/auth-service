package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//Using http.ServeMux is a simple and straightforward
//way to handle HTTP routing in Go. It is part of the
//standard library and is suitable for small applications
//with basic routing needs. However, it lacks advanced
//features and flexibility.

//On the other hand, chi is a more powerful and flexible router
//that provides additional features such as middleware support,
//route grouping, and more. It is well-suited for larger applications
//with more complex routing requirements.

// This is from http inbuilt package from go
func (app *app) mount() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", app.healthCheckHandler)
	return mux
}

// using chi
func (app *app) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", app.healthCheckHandler)
	return r
}
