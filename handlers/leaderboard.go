package handlers

import (
	"github.com/gin-gonic/gin"
	"leaderboard-api/models"
	"leaderboard-api/services"
)

type LeaderboardHandler struct {
	serv *services.LeaderboardService
}

func NewLeaderboardHandler(service *services.LeaderboardService) *LeaderboardHandler {
	return &LeaderboardHandler{
		serv: service,
	}
}

func (h *LeaderboardHandler) SubmitScore(c *gin.Context){
	var score models.Score

	// bind the json to the score object, and return directly if there is an error
	if err := c.BindJSON(&score); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
			"error": err.Error(),
		})

		return
	}

	h.serv.SubmitScore(score)

	c.JSON(200, gin.H{
		"message": "Score submitted successfully",
	})
}

func (h *LeaderboardHandler) GetLeaderboard(c *gin.Context){
	data := h.serv.GetScores()
	c.JSON(200, gin.H{
		"message": "Leaderboard retrieved successfully",
		"data": data,
	})
}

func (h *LeaderboardHandler) Test(c *gin.Context){
	result := h.serv.Test()
	c.JSON(200, gin.H{
		"message": result,
	})
}

func (h *LeaderboardHandler) GetAll(c *gin.Context){
	data := h.serv.GetScores()
	c.JSON(200, gin.H{
		"message": "All scores retrieved successfully",
		"data": data,
	})
}