// Define all the structs here so that it becomes easy to understand.

package main

import "auth-service-2.0/internal/repository"

type app struct {
	AppAddr config
	Store   repository.Storage
}

type config struct {
	AppAddr  string
	DBConfig dbConfig
}

type dbConfig struct {
	DBAddr       string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}
