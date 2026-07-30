package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectToDatabase() *pgx.Conn{
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	conn, err := pgx.Connect(
		context.Background(),
		connString,
	)

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil
	}

	fmt.Println("Connected to Postgre Database")

	return conn
}
