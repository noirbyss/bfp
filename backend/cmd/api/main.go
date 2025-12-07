package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/noirbyss/bfp/internal/database"
	"github.com/noirbyss/bfp/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	errConnectionDataBase = errors.New("failed to connect to the database")
	errCreatingUser       = errors.New("failed to creating user")
)

func createDBConnection() *pgx.Conn {
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(errConnectionDataBase, err)
	}

	return db
}

func printUser(user models.User) {
	fmt.Printf(`
	ID: %d
	Email: %s
	UserName: %s
	Password: %s,
	Created: %s
	`, user.ID, user.Email, user.Username, user.PasswordHash, user.CreatedAt)
}

func generateHashPass(cost int, passwortd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(passwortd, cost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal(errConnectionDataBase, err)
	}

	user := models.User{
		Email:        "testjnjnuj3@gmail.com",
		Username:     "Test3ewrwrwer",
		PasswordHash: generateHashPass(10, []byte("12345678")),
	}

	db := database.UserRepository{
		Db: createDBConnection(),
	}

	if err := db.Create(context.Background(), &user); err != nil {
		log.Fatal(errCreatingUser, err)
	}

	printUser(user)
}
