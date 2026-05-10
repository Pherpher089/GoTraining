package store

import (
	"github.com/Pherpher089/go-tasks/models"
)

type TaskStore interface {
	Add(task models.Task) models.Task
	GetAll() []models.Task
	GetByID(id int) (*models.Task, bool)
	Update(id int, task models.Task) (*models.Task, bool)
	Delete(id int) bool
}