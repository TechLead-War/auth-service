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
		GetById(context.Context, int64) (*Logs, error)
	}
	Sessions interface {
		FetchSessionByID(context.Context, int64) (*Session, error)
		Update(ctx context.Context, session *Session) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users:     &UserStore{Db: db},
		LoginLogs: &LoginLogsStore{Db: db},
		Sessions:  &SessionStore{Db: db},
	}
}
