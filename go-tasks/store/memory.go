package store

import (
	"github.com/Pherpher089/go-tasks/models"
)

type MemoryStore struct {
	tasks []models.Task
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{nextID: 1}
}

func (s *MemoryStore) Add(t models.Task) (models.Task, error) {
		if t.Title == "" {
		return t, ErrEmptyTitle
	}
	t.ID = s.nextID
	s.nextID++
	s.tasks = append(s.tasks, t)
	return t, nil
}
func (s *MemoryStore) GetAll() []models.Task { 
	return s.tasks
}
func (s *MemoryStore) GetByID(id int) (*models.Task, error) { 
	for i := range s.tasks { if s.tasks[i].ID == id { return &s.tasks[i], nil}}
	return nil, &NotFoundError{ID: id}
}
func (s *MemoryStore) Delete(id int) error {	
	tasks := []models.Task{}
	beginingLength := len(s.tasks)
	for _, val := range s.tasks { 
		if val.ID != id { 
			tasks = append(tasks, val)
		}
	}
	s.tasks = tasks
	if beginingLength == len(s.tasks) {
		return &NotFoundError{ID: id}
	} 
	return nil
 }

 func (t *MemoryStore) Update(id int, task models.Task) (models.Task, error) {
	if task.Title == "" {
		return task, &ValidationError{Field: "title", Message: "cannot be empty"}
	}
	taskToUpdate, err := t.GetByID(id)
	if err != nil {
		return task, &NotFoundError{ID: id}
	}
	task.ID = id
	*taskToUpdate = task
	return *taskToUpdate, nil
 }
