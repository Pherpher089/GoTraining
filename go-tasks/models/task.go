package models

import (
	"fmt"
	"time"
)

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

func (t *Task) Complete() { t.Done = true}
func (t Task) String() string { return fmt.Sprintf("[%d] %s", t.ID, t.Title)}