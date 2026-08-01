package services

import (
	"leaderboard-api/models"
	"leaderboard-api/repository"
)

type LeaderboardService struct {
	repo *repository.LeaderboardRepository
}

func NewLeaderboardService(repository *repository.LeaderboardRepository) *LeaderboardService {
	return &LeaderboardService{
		repo: repository,
	}
}

func (s *LeaderboardService) SubmitScore(score models.Score) {
	if (score.Player == "" || score.Score <= 0) {
		return
	}
	s.repo.Save(score)
}

func (s *LeaderboardService) GetScores() []models.Score {
	scores, err := s.repo.GetAll()
	if err != nil {
		return nil
	}
	return scores
}

func (s *LeaderboardService) Test() string {
	s.repo.Test()
	return "Hello World"
}