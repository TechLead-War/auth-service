package repository

import (
	"context"
	"database/sql"
)

type Storage struct {
	Users interface {
		Create(context.Context, *User) error
	}
	LoginLogs interface {
		Create(context.Context, *Logs) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users:     &UsersStore{db: db},
		LoginLogs: &LoginLogsStore{db: db},
	}
}
