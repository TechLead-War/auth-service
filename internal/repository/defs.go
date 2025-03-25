package repository

import (
	"database/sql"
	"time"
)

type Logs struct { // marshal: how these model columns will be named in DB.
	LogID     int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	IPAddress string    `json:"ip_address"`
	Tags      []string  `json:"tags"` // Some metadata about the login, request id, actual log, and others
	CreatedAt time.Time `json:"created_at"`
}

type LoginLogsStore struct {
	Db *sql.DB
}

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"-"` // We will not return password whenever we will marshal, or un-marshal
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
}

type UserStore struct {
	Db *sql.DB
}

type Session struct {
	SessionID int64  `json:"session_id"`
	UserID    int64  `json:"user_id"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
	ExpiresAt string `json:"expires_at"`
	Logs      []Logs `json:"logs"`
	Version   int    `json:"version"` // This is used for optimistic locking, to handle data raise condition
}

type SessionStore struct {
	Db *sql.DB
}
