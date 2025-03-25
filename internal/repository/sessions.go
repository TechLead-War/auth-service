package repository

import (
	"context"
	"database/sql"
	"errors"

	"auth-service-2.0/internal/resources"
	"github.com/lib/pq"
)

func (s *SessionStore) FetchSessionByID(ctx context.Context, sessionId int64) (*Session, error) {
	// Get session details including user ID
	sessionQuery := `
		SELECT s.session_id, s.user_id, s.token, s.created_at, s.expires_at
		FROM sessions s
		WHERE s.session_id = $1
	`
	var session Session
	err := s.Db.QueryRowContext(ctx, sessionQuery, sessionId).Scan(
		&session.SessionID,
		&session.UserID,
		&session.Token,
		&session.CreatedAt,
		&session.ExpiresAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, resources.ErrLogsNotFound
		}
		return nil, err
	}

	logsQuery := `
		SELECT log_id, user_id, ip_address, tags, created_at 
		FROM login_logs
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	// Set a timeout for the query
	ctx, cancel := context.WithTimeout(ctx, resources.QueryTimeOut)
	defer cancel()

	rows, err := s.Db.QueryContext(ctx, logsQuery, session.UserID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var logs []Logs
	for rows.Next() {
		var log Logs
		err := rows.Scan(
			&log.LogID,
			&log.UserID,
			&log.IPAddress,
			pq.Array(&log.Tags),
			&log.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	session.Logs = logs
	return &session, nil
}

func (s *SessionStore) Update(ctx context.Context, session *Session) error {
	query := `
		UPDATE sessions
		SET user_id = $1, token = $2, created_at = $3, expires_at = $4, version = version + 1
		WHERE session_id = $5
		RETURNING version
	`

	// Set a timeout for the query
	ctx, cancel := context.WithTimeout(ctx, resources.QueryTimeOut)
	defer cancel()

	err := s.Db.QueryRowContext(
		ctx,
		query,
		session.UserID,
		session.Token,
		session.CreatedAt,
		session.ExpiresAt,
		session.SessionID,
	).Scan(&session.Version)

	return err
}

func (s *SessionStore) Create(ctx context.Context, session *Session) error {
	query := `
		INSERT INTO sessions (user_id, token, created_at, expires_at, version)
		VALUES ($1, $2, $3, $4, 1)
		RETURNING session_id
	`

	ctx, cancel := context.WithTimeout(ctx, resources.QueryTimeOut)
	defer cancel()

	return s.Db.QueryRowContext(
		ctx,
		query,
		session.UserID,
		session.Token,
		session.CreatedAt,
		session.ExpiresAt,
	).Scan(&session.SessionID)
}
