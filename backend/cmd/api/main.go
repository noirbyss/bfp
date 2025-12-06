package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var (
	errConnectionDataBase = errors.New("failed to connect to the database")
)

func main() {
	if err := godotenv.Load("../../internal/config/.env"); err != nil {
		log.Fatal("Error load .env file", err)
	}

	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(errConnectionDataBase)
	}
	defer db.Close(context.Background())

	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(errConnectionDataBase)
	}
	fmt.Println("Connected")
}
