package handlers

import (
	"net/http"

	"github.com/Pherpher089/go-tasks/store"
)

type TaskHandler struct {
	store store.TaskStore
}

func (t *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {

} 

func (t *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {

}

func (h *TaskHandler) HanldeTasks(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		h.ListTasks(w,r)
	case http.MethodPost:
		h.CreateTask(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}