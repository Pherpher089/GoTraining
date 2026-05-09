package store

import (
	"github.com/Pherpher089/go-tasks/models"
)

type MemoryStore struct {
	tasks []models.Task
	nextID int
}

func (s *MemoryStore) Add(t models.Task) models.Task { 
	s.tasks = append(s.tasks, t)
	return t
}
func (s *MemoryStore) GetAll() []models.Task { 
	return s.tasks
}
func (s *MemoryStore) GetByID(id int) (*models.Task, bool) { 
	for _, val := range s.tasks { if val.ID == id { return &val, true}}
	return nil, false
}
func (s *MemoryStore) Delete(id int) bool {	
	tasks := []models.Task{}
	beginingLength := len(s.tasks)
	for _, val := range s.tasks { 
		if val.ID != id { 
			tasks = append(tasks, val)
		}
	}
	if beginingLength == len(s.tasks) {
		return true
	} else {
		return false
	}
 }