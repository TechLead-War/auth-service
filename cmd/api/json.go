package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// If we want all the required fields to be passed on the payload, we can use the validator package.
var Validate *validator.Validate

// Each time service starts it will be called.
func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func (app *app) JsonResponse(w http.ResponseWriter, status int, data any) error {
	return WriteJson(w, status, &ResponseEnvelope{Data: data})
}

func WriteJsonError(w http.ResponseWriter, r *http.Request, status int, err error) error {

	log.Printf("Error: %s Method: %s Path: %s", err.Error(), r.Method, r.URL.Path)
	return WriteJson(w, status, &ErrorEnvelope{Error: err.Error()})
}

func ReadJson(r *http.Request, data any) error {
	r.Body = http.MaxBytesReader(nil, r.Body, maxJsonBodySize) // max size of JSON body

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // disallow any other structs that we don't have

	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	return nil
}
