package repository

import (
	"context"
	"database/sql"

	"auth-service-2.0/internal/resources"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

func (lls *LoginLogsStore) Create(ctx context.Context, log *Logs) error {
	query := `
		INSERT INTO login_logs 
		    (user_id, ip_address, tags, created_at) 
		VALUES 
		    ($1, $2, $3, $4)  
		RETURNING 
			log_id, created_at
	`

	err := lls.Db.QueryRowContext(
		ctx, query,

		// fill values in $1, $2, $3, $4
		log.UserID,
		log.IPAddress,
		pq.Array(log.Tags),
		log.CreatedAt,
	).Scan(
		&log.LogID,
		&log.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (lls *LoginLogsStore) GetById(ctx context.Context, log_id int64) (*Logs, error) {
	query := `
		SELECT user_id, ip_address, tags, created_at 
		FROM login_logs 
		WHERE log_id = $1 ;
	`

	// Set a timeout for the query
	ctx, cancel := context.WithTimeout(ctx, resources.QueryTimeOut)
	defer cancel()

	var log Logs
	err := lls.Db.QueryRowContext(ctx, query, log_id).Scan(
		&log.UserID,
		&log.IPAddress,
		pq.Array(&log.Tags),
		&log.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, resources.ErrLogsNotFound
		default:
			return nil, err
		}
	}
	return &log, nil
}
