package models

import "time"

type Priority string

const (
	Low 	Priority = "low"
	Medium	Priority = "medium"
	High	Priority = "high"
)

type Task struct {
	ID			int			`json:"id"`
	Title		string		`json:"title"`
	Description	string		`json:"description"`
	Done		bool		`json:"done"`
	Priority	Priority	`json:"priority"`
	CreatedAt	time.Time	`json:"created_at"`
}