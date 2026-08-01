package main

import (
	"leaderboard-api/database"
	"leaderboard-api/handlers"
	"leaderboard-api/repository"
	"leaderboard-api/services"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	_ = godotenv.Load()

	conn := database.ConnectToDatabase()
	
	defer conn.Close()

	repo := repository.NewLeaderboardRepository(conn)
	service := services.NewLeaderboardService(repo)
	handlers := handlers.NewLeaderboardHandler(service)

	r := gin.Default()

	r.POST("/scores", handlers.SubmitScore)
	r.GET("/leaderboard", handlers.GetLeaderboard)
	r.GET("/all", handlers.GetAll)
	r.GET("/", handlers.Test)

	slog.Info("Server is running on port 8080")
	r.Run(":8080")
}