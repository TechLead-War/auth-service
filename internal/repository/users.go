package repository

import (
	"context"

	"auth-service-2.0/internal/resources"
)

func (us *UserStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users 
		    (username, password, created_at, email) 
		VALUES 
		    ($1, $2, $3, $4)  
		RETURNING 
			id, created_at
	`

	// Set a timeout for the query
	ctx, cancel := context.WithTimeout(ctx, resources.QueryTimeOut)
	defer cancel()

	err := us.Db.QueryRowContext(
		ctx, query,
		user.Username,
		user.Password,
		user.CreatedAt,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}
