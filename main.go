package main

import (
	"context"
	// "fmt"
	"log/slog"
	"leaderboard-api/database"
	// "leaderboard-api/handlers"

	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	_ = godotenv.Load()

	// r := gin.Default()

	// r.POST("/scores", handlers.SubmitScore)
	// r.GET("/leaderboard", handlers.GetLeaderboard)
	// r.GET("/", handlers.Test)

	// r.Run(":8080")

	conn := database.ConnectToDatabase()
	
	defer conn.Close(context.Background())

	var version string
	var databaseTest int

	err:= conn.QueryRow(
		context.Background(), 
		"SELECT version()",
	).Scan(&version)

	if err != nil {
		panic(err)
	}

	err2 := conn.QueryRow(
		context.Background(), 
		"SELECT gold from turn_snapshots limit 1",
	).Scan(&databaseTest)

	if err2 != nil {
		panic(err2)
	}

	slog.Info("Database version:", version)
	slog.Info("Database test value:", databaseTest)
}