// Day 6 Exercise: JSON & HTTP Handlers
//
// Run this file with: go run exercises/day6_http_handlers.go
// Then test it in a second terminal with curl (examples shown per section).
// Press Ctrl+C to stop the server.
//
// This exercise is different from previous ones — instead of seeing output
// immediately, you'll start a server and hit it with curl to verify your work.
//
// Work through each TODO. Build up the handler step by step.
// Uncomment the answers at the bottom when you're done.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// =============================================================================
// Types (same as previous days, self-contained for this exercise)
// =============================================================================

type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

type Task struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Done     bool     `json:"done"`
	Priority Priority `json:"priority"`
}

// JSON struct tags tell encoding/json what field names to use in JSON output.
// Without them, Go would use the exact field name (e.g. "ID" instead of "id").

// ----- Errors -----

var ErrEmptyTitle = errors.New("task title cannot be empty")

type NotFoundError struct{ ID int }

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("task with ID %d not found", e.ID)
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}

// ----- Store -----

type TaskStore interface {
	Add(task Task) (Task, error)
	GetAll() []Task
	GetByID(id int) (*Task, error)
	Delete(id int) error
	Update(id int, task Task) (Task, error)
}

type MemoryStore struct {
	tasks  []Task
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{nextID: 1}
}

func (s *MemoryStore) Add(t Task) (Task, error) {
	if t.Title == "" {
		return Task{}, &ValidationError{Field: "title", Message: "cannot be empty"}
	}
	t.ID = s.nextID
	s.nextID++
	s.tasks = append(s.tasks, t)
	return t, nil
}

func (s *MemoryStore) GetAll() []Task {
	return s.tasks
}

func (s *MemoryStore) GetByID(id int) (*Task, error) {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			return &s.tasks[i], nil
		}
	}
	return nil, &NotFoundError{ID: id}
}

func (s *MemoryStore) Delete(id int) error {
	var updated []Task
	found := false
	for _, t := range s.tasks {
		if t.ID == id {
			found = true
			continue
		}
		updated = append(updated, t)
	}
	if !found {
		return &NotFoundError{ID: id}
	}
	s.tasks = updated
	return nil
}

func (t *MemoryStore) Update(id int, task Task) (Task, error) {
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

// =============================================================================
// PART 1: JSON encoding and decoding basics
// =============================================================================
// Before writing handlers, understand how Go converts between structs and JSON.
//
// Encoding: struct → JSON bytes   json.NewEncoder(w).Encode(v)
// Decoding: JSON bytes → struct   json.NewDecoder(r.Body).Decode(&v)

// TODO: Write encodeJSON(w http.ResponseWriter, status int, v any)
//       It should:
//       1. Set the Content-Type header to "application/json"
//       2. Write the status code with w.WriteHeader(status)
//       3. Encode v to w using json.NewEncoder
//       Return any encoding error.

func encodeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)

}


// TODO: Write decodeJSON(r *http.Request, v any) error
//       Decode the request body into v using json.NewDecoder.
//       Return any decoding error.

func decodeJSON(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// =============================================================================
// PART 2: Error → HTTP status code mapping
// =============================================================================
// Handlers shouldn't return errors directly — they need to send HTTP responses.
// This helper maps your store error types to the right HTTP status codes.

// TODO: Write handleError(w http.ResponseWriter, err error)
//       Use errors.As to check the error type and respond accordingly:
//       - *NotFoundError  → 404, json body: {"error": "task with ID X not found"}
//       - *ValidationError → 400, json body: {"error": "validation error on ..."}
//       - anything else   → 500, json body: {"error": "internal server error"}
//
//       Use encodeJSON to send the response. For the body, use a
//       map[string]string{"error": err.Error()}

func handleError(w http.ResponseWriter, err error) {
	nfe := &NotFoundError{}
	ve := &ValidationError{}

	switch {
	case errors.As(err, &nfe): 
		encodeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
	case errors.As(err, &ve):
		encodeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
	default:
		encodeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal service error"})
	}
}


// =============================================================================
// PART 3: The TaskHandler struct
// =============================================================================
// Rather than using global state, we attach the store to a handler struct.
// Each method on TaskHandler is an http.HandlerFunc.

type TaskHandler struct {
	store TaskStore
}

func NewTaskHandler(store TaskStore) *TaskHandler {
	return &TaskHandler{store: store}
}

// TODO: Write HandleTasks(w http.ResponseWriter, r *http.Request)
//       This is the entry point for /tasks — route by HTTP method:
//       GET  → h.listTasks(w, r)
//       POST → h.createTask(w, r)
//       anything else → 405 Method Not Allowed

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

// TODO: Write HandleTask(w http.ResponseWriter, r *http.Request)
//       Entry point for /tasks/{id} — extract the ID from the URL path,
//       then route by method:
//       GET    → h.getTask(w, r, id)
//       DELETE → h.deleteTask(w, r, id)
//       anything else → 405
//

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
// =============================================================================
// PART 4: The individual endpoint handlers
// =============================================================================

// TODO: Write listTasks — GET /tasks
//       Get all tasks from the store and encode them as JSON with status 200.
//       If the store is empty, return an empty JSON array [] not null.
//       Hint: if tasks is nil, set it to []Task{}

func (t *TaskHandler) listTasks(w http.ResponseWriter, r *http.Request) {
	tasks := t.store.GetAll()
	if tasks == nil {
		tasks = []Task{}
	}
	encodeJSON(w, http.StatusOK, tasks)
}

// TODO: Write createTask — POST /tasks
//       Decode the request body into a Task.
//       If decoding fails → 400 {"error": "invalid request body"}
//       Add to the store — if that fails, call handleError.
//       On success → 201 Created with the created task as JSON.

func(t *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
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

// TODO: Write getTask — GET /tasks/{id}
//       Look up the task by ID.
//       If not found, call handleError (it'll send 404).
//       On success → 200 with the task as JSON.

func (t *TaskHandler) getTask(w http.ResponseWriter, r *http.Request, id int) {
	task, err := t.store.GetByID(id)
	if err != nil {
		handleError(w, err)
		return
	}
	encodeJSON(w, http.StatusOK, task)
}

// TODO: Write deleteTask — DELETE /tasks/{id}
//       Delete by ID. If not found, call handleError.
//       On success → 204 No Content (no body).

func (t *TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request, id int) {
	err := t.store.Delete(id)
	if err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// =============================================================================
// PART 5: Wiring it all together in main
// =============================================================================

func main() {
	store := NewMemoryStore()
	handler := NewTaskHandler(store)

	mux := http.NewServeMux()

	// TODO: Register two routes:
	//   /tasks    → handler.HandleTasks
	//   /tasks/   → handler.HandleTask  (trailing slash catches /tasks/1, /tasks/2 etc)
	//

	mux.HandleFunc("/tasks", handler.HandleTasks)
	mux.HandleFunc("/tasks/", handler.HandleTask)


	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

	// Suppress unused import warnings while parts are commented out
	_ = strings.Split
	_ = strconv.Atoi
	_ = json.NewEncoder
}

func (t *TaskHandler) update(w http.ResponseWriter, r *http.Request, id int) {
	var task Task
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

// =============================================================================
// TEST WITH CURL
// =============================================================================
// Open a second terminal and run these commands to verify each endpoint.
// The server must be running first: go run exercises/day6_http_handlers.go
//
// List tasks (empty):
//   curl http://localhost:8080/tasks
//   Expected: []
//
// Create a task:
//   curl -X POST http://localhost:8080/tasks \
//     -H "Content-Type: application/json" \
//     -d '{"title":"Learn Go","priority":"high"}'
//   Expected: {"id":1,"title":"Learn Go","done":false,"priority":"high"}
//
// Create another:
//   curl -X POST http://localhost:8080/tasks \
//     -H "Content-Type: application/json" \
//     -d '{"title":"Write tests","priority":"medium"}'
//
// List all tasks:
//   curl http://localhost:8080/tasks
//   Expected: [{"id":1,...},{"id":2,...}]
//
// Get one task:
//   curl http://localhost:8080/tasks/1
//   Expected: {"id":1,"title":"Learn Go","done":false,"priority":"high"}
//
// Get missing task:
//   curl http://localhost:8080/tasks/99
//   Expected: {"error":"task with ID 99 not found"}  (404)
//
// Create with empty title (validation error):
//   curl -X POST http://localhost:8080/tasks \
//     -H "Content-Type: application/json" \
//     -d '{"title":""}'
//   Expected: {"error":"validation error on title: cannot be empty"}  (400)
//
// Delete a task:
//   curl -X DELETE http://localhost:8080/tasks/1
//   Expected: (empty body, 204 status)
//
// Confirm it's gone:
//   curl http://localhost:8080/tasks/1
//   Expected: {"error":"task with ID 1 not found"}  (404)

// =============================================================================
// PART 6: Challenge — PUT /tasks/{id} (update endpoint)
// =============================================================================
// Add an Update(id int, task Task) (Task, error) method to MemoryStore.
// Add an updateTask(w, r, id) method to TaskHandler.
// Wire it into HandleTask under http.MethodPut.
//
// The update should:
//   - Return NotFoundError if the ID doesn't exist
//   - Return ValidationError if the new title is empty
//   - Replace the task in the slice and return the updated task with 200 OK


//
// Test with curl:
//   curl -X PUT http://localhost:8080/tasks/2 \
//     -H "Content-Type: application/json" \
//     -d '{"title":"Write better tests","priority":"high","done":true}'
//   Expected: {"id":2,"title":"Write better tests","done":true,"priority":"high"}

// =============================================================================
// ANSWERS — uncomment to check your work
// =============================================================================

// func encodeJSON(w http.ResponseWriter, status int, v any) error {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	return json.NewEncoder(w).Encode(v)
// }

// func decodeJSON(r *http.Request, v any) error {
// 	return json.NewDecoder(r.Body).Decode(v)
// }

// func handleError(w http.ResponseWriter, err error) {
// 	var nfe *NotFoundError
// 	var ve *ValidationError
// 	switch {
// 	case errors.As(err, &nfe):
// 		encodeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
// 	case errors.As(err, &ve):
// 		encodeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
// 	default:
// 		encodeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
// 	}
// }

// func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		h.listTasks(w, r)
// 	case http.MethodPost:
// 		h.createTask(w, r)
// 	default:
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

// func (h *TaskHandler) HandleTask(w http.ResponseWriter, r *http.Request) {
// 	parts := strings.Split(r.URL.Path, "/")
// 	id, err := strconv.Atoi(parts[len(parts)-1])
// 	if err != nil {
// 		encodeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid task ID"})
// 		return
// 	}
// 	switch r.Method {
// 	case http.MethodGet:
// 		h.getTask(w, r, id)
// 	case http.MethodDelete:
// 		h.deleteTask(w, r, id)
// 	default:
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

// func (h *TaskHandler) listTasks(w http.ResponseWriter, r *http.Request) {
// 	tasks := h.store.GetAll()
// 	if tasks == nil {
// 		tasks = []Task{}
// 	}
// 	encodeJSON(w, http.StatusOK, tasks)
// }

// func (h *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
// 	var task Task
// 	if err := decodeJSON(r, &task); err != nil {
// 		encodeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
// 		return
// 	}
// 	created, err := h.store.Add(task)
// 	if err != nil {
// 		handleError(w, err)
// 		return
// 	}
// 	encodeJSON(w, http.StatusCreated, created)
// }

// func (h *TaskHandler) getTask(w http.ResponseWriter, r *http.Request, id int) {
// 	task, err := h.store.GetByID(id)
// 	if err != nil {
// 		handleError(w, err)
// 		return
// 	}
// 	encodeJSON(w, http.StatusOK, task)
// }

// func (h *TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request, id int) {
// 	if err := h.store.Delete(id); err != nil {
// 		handleError(w, err)
// 		return
// 	}
// 	w.WriteHeader(http.StatusNoContent)
// }
