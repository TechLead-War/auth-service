// Define all the structs here so that it becomes easy to understand.

package main

import (
	"time"

	"auth-service-2.0/internal/repository"
)

type app struct {
	AppAddr config
	Store   repository.Storage
}

type config struct {
	AppAddr  string
	DBConfig dbConfig
	Env      string
}

type dbConfig struct {
	DBAddr       string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

type CreateLogPayload struct {
	Log              string    `json:"log" validate:"required,max=10000"`
	RequestId        string    `json:"request_id" validate:"required"`
	RequestTimeStamp time.Time `json:"time_stamp"`
	UserId           int64     `json:"user_id" validate:"required,gt=0"`
}

type ErrorEnvelope struct {
	Error string `json:"error"`
}

type ResponseEnvelope struct {
	Data any `json:"data"`
}
