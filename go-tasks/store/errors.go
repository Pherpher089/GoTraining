package store

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("task not found")
var ErrEmptyTitle = errors.New("task title cannot be empty")
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}

type NotFoundError struct {
	ID int
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("task with ID %d not found", n.ID)
}

