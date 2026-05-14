package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Pherpher089/go-tasks/models"
	"github.com/Pherpher089/go-tasks/store"
)
type TaskHandler struct {
	store store.TaskStore
}

func handleError(w http.ResponseWriter, err error) {
	nfe := &store.NotFoundError{}
	ve := &store.ValidationError{}

	switch {
	case errors.As(err, &nfe): 
		encodeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	case errors.As(err, &ve):
		encodeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	default:
		encodeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal service error"})
	}
}

func NewTaskHandler(store store.TaskStore) *TaskHandler {
	return &TaskHandler{store: store}
}


func (t *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.MethodGet:
		t.listTasks(w, r)
	case http.MethodPost:
		t.createTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		
	}
}

func encodeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)

}

func decodeJSON(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (t *TaskHandler) HandleTask(w http.ResponseWriter, r *http.Request){
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])

	if err != nil {
		encodeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid task"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		t.getTask(w, r, id)
	case http.MethodDelete:
		t.deleteTask(w, r, id)
	case http.MethodPut:
		t.update(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func (t *TaskHandler) listTasks(w http.ResponseWriter, r *http.Request) {
	tasks := t.store.GetAll()
	if tasks == nil {
		tasks = []models.Task{}
	}
	encodeJSON(w, http.StatusOK, tasks)
}

func(t *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := decodeJSON(r, &task); err != nil {
		encodeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return 
	}

	created, err := t.store.Add(task)
	if err != nil {
		handleError(w, err)
		return
	}
	encodeJSON(w, http.StatusCreated, created)
}

func (t *TaskHandler) getTask(w http.ResponseWriter, r *http.Request, id int) {
	task, err := t.store.GetByID(id)
	if err != nil {
		handleError(w, err)
		return
	}
	encodeJSON(w, http.StatusOK, task)
}

func (t *TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request, id int) {
	err := t.store.Delete(id)
	if err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (t *TaskHandler) update(w http.ResponseWriter, r *http.Request, id int) {
	var task models.Task
	if err := decodeJSON(r, &task); err != nil {
        encodeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
        return
    }

	updated, err := t.store.Update(id, task)
    if err != nil {
        handleError(w, err)
        return
    }

	encodeJSON(w, http.StatusOK, updated)
}