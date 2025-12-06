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
	ctx                   = context.Background()
)

func main() {
	sqlRequest, err := os.ReadFile("../migrations/001_create_users_table.sql")
	if err != nil {
		log.Fatal("error reading migration file", err)
	}

	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(errConnectionDataBase)
	}

	db, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(errConnectionDataBase)
	}
	defer db.Close(ctx)

	if _, err := db.Exec(ctx, string(sqlRequest)); err != nil {
		log.Fatal("error request")
	}

	fmt.Println("Migration complete!")
}
