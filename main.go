package main

import (
	"leaderboard-api/handlers"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.POST("/scores", handlers.SubmitScore)
	r.GET("/leaderboard", handlers.GetLeaderboard)
	r.GET("/", handlers.Test)

	r.Run(":8080")
}