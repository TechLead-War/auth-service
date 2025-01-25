package repository

import (
	"context"

	"github.com/lib/pq"
)

func (lls *LoginLogsStore) Create(ctx context.Context, log *Logs) error {
	query := `
		INSERT INTO login_logs 
		    (user_id, ip_address, tags, created_at, updated_at) 
		VALUES 
		    ($1, $2, $3, $4, $5)  
		RETURNING 
			id, created_at, updated_at
	`

	err := lls.db.QueryRowContext(
		ctx, query,
		log.UserID,
		log.IPAddress,
		pq.Array(log.Tags),
		log.CreatedAt,
		log.UpdateAt,
	).Scan(
		&log.ID,
		&log.CreatedAt,
		&log.UpdateAt)

	if err != nil {
		return err
	}

	return nil
}
