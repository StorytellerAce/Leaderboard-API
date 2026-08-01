package models

type Score struct {
	ID		int    `json:"id"`
	Player	string `json:"player"`
	Score	int    `json:"score"`
}