package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/noirbyss/bfp/internal/models"
)

type UserRepository struct {
	Db *pgx.Conn
}

// Create adds a user to the database.
func (database UserRepository) Create(ctx context.Context, user *models.User) error {
	request := `INSERT INTO users (email, username, password_hash)
	VALUES ($1, $2, $3) RETURNING id, created_at`

	err := database.Db.QueryRow(ctx, request, user.Email, user.Username, user.PasswordHash).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
