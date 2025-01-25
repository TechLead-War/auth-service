package repository

import (
	"context"
)

func (us *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users 
		    (username, password, created_at, email) 
		VALUES 
		    ($1, $2, $3, $4)  
		RETURNING 
			id, created_at
	`

	err := us.db.QueryRowContext(
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
