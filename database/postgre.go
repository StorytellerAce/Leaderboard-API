package database

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDatabase() *pgxpool.Pool{
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	pool, err := pgxpool.New(
		context.Background(),
		connString,
	)

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil
	}

	fmt.Println("Connected to Postgre Database")

	return pool
}
