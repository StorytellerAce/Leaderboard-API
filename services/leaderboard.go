package services

import (
	"leaderboard-api/models"
	"leaderboard-api/repository"
)

func SubmitScore(score models.Score) {
	if (score.Player == "" || score.Score <= 0) {
		return
	}
	repository.Save(score)
}

func GetScores() []models.Score {
	return repository.GetAll()
}

func Test() string {
	return "Hello World"
}