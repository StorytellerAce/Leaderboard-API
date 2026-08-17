package repository

import (
	"context"
	"leaderboard-api/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

//constructor
type LeaderboardRepository struct {
	db *pgxpool.Pool
}

func NewLeaderboardRepository(
	db *pgxpool.Pool,
) *LeaderboardRepository {
	return &LeaderboardRepository{
		db: db,
	}
}

func (r *LeaderboardRepository) Save(score models.Score){
	_, err := r.db.Exec(
		context.Background(),
		"INSERT INTO scores (player, score) VALUES ($1, $2)",
		score.Player,
		score.Score,
	)
	if err != nil {
		slog.Error("Error saving score", "error", err)
	}
}

func (r *LeaderboardRepository) GetAll() ([]models.Score, error) {
	rows, err := r.db.Query(
		context.Background(),
		"SELECT id, player, score FROM scores ORDER BY score DESC",
	)

	if err != nil {
		slog.Error("Error fetching scores", "error", err)
		return nil, err
	}
	defer rows.Close()

	var scores []models.Score
	for rows.Next(){
		var score models.Score
		err := rows.Scan(
			&score.ID,
			&score.Player,
			&score.Score,
		)
		if err != nil {
			slog.Error("Error scanning score", "error", err)
			return nil, err
		}
		scores = append(scores, score)
	}
	return scores, nil
}

func (r *LeaderboardRepository) Test() {
	var version string
	var databaseTest int

	err:= r.db.QueryRow(
		context.Background(), 
		"SELECT version()",
	).Scan(&version)

	if err != nil {
		panic(err)
	}

	err2 := r.db.QueryRow(
		context.Background(), 
		"SELECT gold from turn_snapshots limit 1",
	).Scan(&databaseTest)

	if err2 != nil {
		panic(err2)
	}

	slog.Info("Database version:", version)
	slog.Info("Database test value:", databaseTest)
}