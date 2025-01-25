package repository

import "database/sql"

type Logs struct { // marshal: how these model columns will be named in DB.
	ID        int64    `json:"id"`
	UserID    int64    `json:"user_id"`
	IPAddress string   `json:"ip_address"`
	Tags      []string `json:"tags"` // Some metadata about the login
	CreatedAt string   `json:"created_at"`
	UpdateAt  string   `json:"updated_at"`
}

type LoginLogsStore struct {
	db *sql.DB
}

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"-"` // We will not return password whenever we will marshal, or un-marshal
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
}

type UsersStore struct {
	db *sql.DB
}
