package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/noirbyss/bfp/internal/models"
)

type UserRepository struct {
	Db *pgx.Conn
}

func (ur UserRepository) Create(ctx context.Context, user *models.User) error {
	request := `INSERT INTO users (email, username, password_hash)
	VALUES ($1, $2, $3) RETURNING id, created_at`

	err := ur.Db.QueryRow(ctx, request, user.Email, user.Username, user.Password_Hash).Scan(&user.Id, &user.CreatedAt)
	if err != nil {
		return err
	}
	log.Println("User created!")
	fmt.Printf("User id: %d\nCreated: %s", user.Id, user.CreatedAt)

	return nil
}
