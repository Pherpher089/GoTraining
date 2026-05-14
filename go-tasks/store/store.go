package store

import (
	"github.com/Pherpher089/go-tasks/models"
)
type TaskStore interface {
    Add(task models.Task) (models.Task, error)
    GetAll() []models.Task
    GetByID(id int) (*models.Task, error)
    Update(id int, task models.Task) (models.Task, error)
    Delete(id int) error
}