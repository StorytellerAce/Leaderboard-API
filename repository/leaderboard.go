package repository

import "leaderboard-api/models"

var scores []models.Score

func Save(score models.Score){
	scores = append(scores, score)
}

func GetAll() []models.Score{
	return scores
}