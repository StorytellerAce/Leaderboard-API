package handlers

import (
	"github.com/gin-gonic/gin"
	"leaderboard-api/models"
	"leaderboard-api/services"
)

func SubmitScore(c *gin.Context){
	var score models.Score

	// bind the json to the score object, and return directly if there is an error
	if err := c.BindJSON(&score); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
			"error": err.Error(),
		})

		return
	}

	services.SubmitScore(score)

	c.JSON(200, gin.H{
		"message": "Score submitted successfully",
	})
}

func GetLeaderboard(c *gin.Context){
	data := services.GetScores()
	c.JSON(200, gin.H{
		"message": "Leaderboard retrieved successfully",
		"data": data,
	})
}

func Test(c *gin.Context){
	result := services.Test()
	c.JSON(200, gin.H{
		"message": result,
	})
}

func TestDatabaseConnection(c *gin.Context){
	
}