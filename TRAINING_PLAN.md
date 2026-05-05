# Go Interview Training Plan
**Goal:** Go proficiency in ~10 sessions (2 hrs/day) ahead of your interview  
**Project:** `go-tasks` — a Task Manager REST API built incrementally  
**Format:** Each session = one new Go concept + one new API feature + a short standalone exercise

---

## The Project: `go-tasks`

A JSON REST API for managing tasks. By the end you'll have:

```
GET    /tasks           → list all tasks
GET    /tasks/{id}      → get one task
POST   /tasks           → create a task
PUT    /tasks/{id}      → update a task
DELETE /tasks/{id}      → delete a task
```

Built across 10 sessions, each session layering in a new Go concept.

---

## Day 1 — Setup, Arrays & Slices
**Concept:** How Go handles ordered collections  
**Build:** Initialize the project and get a "Hello, Go!" HTTP server running  
**Exercise:** `exercises/day1_slices.go`

### Project Init (you do this yourself!)
```bash
# In your GoTraining folder:
mkdir go-tasks
cd go-tasks
go mod init github.com/chris/go-tasks

# Create folders
mkdir handlers models store exercises
```

Then create `main.go`:
```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "go-tasks API is running!")
    })
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Run it:
```bash
go run main.go
# Visit http://localhost:8080 in your browser
```

### Key Concepts
- Arrays have fixed size: `var ids [5]int`
- Slices are dynamic and what you'll use 99% of the time: `tasks := []string{}`
- `append(slice, item)` — always reassign: `tasks = append(tasks, "new task")`
- `range` iterates: `for i, v := range tasks { ... }`
- Slicing a slice: `tasks[1:3]` — includes index 1, excludes 3

---

## Day 2 — Structs & Pointers
**Concept:** Custom types and memory addressing  
**Build:** Define the `Task` struct in `models/task.go`

### What to build
```go
// models/task.go
package models

import "time"

type Priority string

const (
    Low    Priority = "low"
    Medium Priority = "medium"
    High   Priority = "high"
)

type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Done        bool      `json:"done"`
    Priority    Priority  `json:"priority"`
    CreatedAt   time.Time `json:"created_at"`
}
```

### Key Concepts
- Struct literal: `t := Task{Title: "Buy milk", Priority: Medium}`
- Pointer to struct: `t := &Task{...}` — changes inside functions affect the original
- Dereferencing: `*t` to get the value, `t.Title` works directly (Go auto-dereferences)
- When to use pointers:
  - Large structs (avoid copying)
  - When you need to mutate the original
  - When nil is a valid state

### Exercise
```go
// Create a Task, then write a function that takes *Task and marks it done.
// Observe what happens if you take Task (value) instead — the original won't change.
```

---

## Day 3 — Methods
**Concept:** Functions attached to types  
**Build:** Add methods to `Task` and create an in-memory store in `store/memory.go`

### Key Concepts
- Value receiver: `func (t Task) String() string` — gets a copy, can't mutate
- Pointer receiver: `func (t *Task) Complete()` — mutates the original
- Convention: if ANY method uses a pointer receiver, use pointer receivers for ALL methods on that type

### What to build
```go
// On Task:
func (t *Task) Complete() { t.Done = true }
func (t Task) String() string { return fmt.Sprintf("[%d] %s", t.ID, t.Title) }

// In store/memory.go:
type MemoryStore struct {
    tasks  []models.Task
    nextID int
}

func (s *MemoryStore) Add(t models.Task) models.Task { ... }
func (s *MemoryStore) GetAll() []models.Task { ... }
func (s *MemoryStore) GetByID(id int) (*models.Task, bool) { ... }
func (s *MemoryStore) Delete(id int) bool { ... }
```

---

## Day 4 — Interfaces
**Concept:** Behavior contracts — Go's most powerful feature  
**Build:** Define a `TaskStore` interface; wire it into handlers

### Key Concepts
- An interface is a set of method signatures
- A type *implicitly* satisfies an interface if it has all the methods (no `implements` keyword)
- This lets you swap implementations (in-memory → database) without changing handler code
- The empty interface `interface{}` (or `any`) accepts anything

### What to build
```go
// store/store.go
package store

import "github.com/chris/go-tasks/models"

type TaskStore interface {
    Add(task models.Task) models.Task
    GetAll() []models.Task
    GetByID(id int) (*models.Task, bool)
    Update(id int, task models.Task) (*models.Task, bool)
    Delete(id int) bool
}
```

Then update `main.go` to use `store.TaskStore` as the type passed to handlers — not `*store.MemoryStore` directly. This is the repository pattern.

---

## Day 5 — Error Handling
**Concept:** Go's explicit, value-based error system  
**Build:** Add proper errors throughout — custom error types, wrapped errors

### Key Concepts
- `error` is just an interface: `type error interface { Error() string }`
- Return errors as last return value: `func GetByID(id int) (*Task, error)`
- Check errors immediately: `if err != nil { ... }`
- Custom error type: `type NotFoundError struct { ID int }`
- Wrap errors for context: `fmt.Errorf("getByID: %w", err)`
- Unwrap: `errors.Is(err, ErrNotFound)`, `errors.As(err, &target)`

### What to build
```go
// store/errors.go
var ErrNotFound = errors.New("task not found")

type ValidationError struct {
    Field   string
    Message string
}
func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}
```

Update all store methods to return `error`. Update handlers to send proper HTTP status codes (404, 400, 500) based on error type.

---

## Day 6 — JSON & HTTP Handlers
**Concept:** Encoding/decoding JSON, writing real HTTP handlers  
**Build:** Complete all 5 CRUD endpoints

### Key Concepts
- Decode request body: `json.NewDecoder(r.Body).Decode(&task)`
- Encode response: `json.NewEncoder(w).Encode(task)`
- Always set content type: `w.Header().Set("Content-Type", "application/json")`
- Status codes: `w.WriteHeader(http.StatusCreated)` — must be called BEFORE writing body
- `http.ServeMux` for routing (or introduce `chi` router for path params)

### What to build
```go
// handlers/task_handler.go
type TaskHandler struct {
    store store.TaskStore
}

func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        h.listTasks(w, r)
    case http.MethodPost:
        h.createTask(w, r)
    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}
```

Test with curl:
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go","priority":"high"}'

curl http://localhost:8080/tasks
```

---

## Day 7 — Goroutines & Channels
**Concept:** Go's built-in concurrency primitives  
**Build:** Add a background worker that auto-expires overdue tasks

### Key Concepts
- `go func()` launches a goroutine (lightweight thread, ~2KB stack)
- Channels send data between goroutines: `ch := make(chan int)`
- Buffered channels don't block until full: `ch := make(chan int, 10)`
- `select` waits on multiple channels — like a switch for channels
- `sync.WaitGroup` — wait for a group of goroutines to finish
- `sync.Mutex` — protect shared data from concurrent access

### What to build
```go
// A background goroutine that checks for overdue tasks every minute
func startWorker(store store.TaskStore, quit <-chan struct{}) {
    ticker := time.NewTicker(1 * time.Minute)
    defer ticker.Stop()
    for {
        select {
        case <-ticker.C:
            // check and flag overdue tasks
        case <-quit:
            log.Println("worker shutting down")
            return
        }
    }
}

// In main.go:
quit := make(chan struct{})
go startWorker(store, quit)
```

Also add a `sync.RWMutex` to `MemoryStore` so concurrent requests don't cause races:
```go
type MemoryStore struct {
    mu     sync.RWMutex
    tasks  []models.Task
    nextID int
}
```

---

## Day 8 — Maps & Advanced Data Structures
**Concept:** Go's key-value map type and when to choose it over slices  
**Build:** Add tag-based filtering and a simple in-memory index

### Key Concepts
- `m := map[string]int{}` — always initialize before use (nil map panics on write)
- Check existence: `val, ok := m[key]` — `ok` is false if key not present
- Delete: `delete(m, key)`
- Maps are unordered — don't rely on iteration order
- Map of slices: `map[string][]Task` for grouping

### What to build
- Add `Tags []string` to the `Task` struct
- Add `GET /tasks?tag=work` filtering endpoint
- Internally maintain a `map[string][]int` tag index for O(1) tag lookups

---

## Day 9 — Testing
**Concept:** Go's built-in test framework — table-driven tests  
**Build:** Test suite for the store and handlers

### Key Concepts
- Test files end in `_test.go`, test functions start with `Test`
- Run tests: `go test ./...`
- Table-driven tests: define a slice of test cases, loop and run each
- `t.Run("name", func(t *testing.T) {...})` for subtests
- Use `httptest.NewRecorder()` and `httptest.NewRequest()` for handler tests

### What to build
```go
// store/memory_test.go
func TestMemoryStore_Add(t *testing.T) {
    tests := []struct {
        name  string
        input models.Task
        want  int // expected ID
    }{
        {"first task", models.Task{Title: "Buy milk"}, 1},
        {"second task", models.Task{Title: "Walk dog"}, 2},
    }
    s := &MemoryStore{}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := s.Add(tt.input)
            if got.ID != tt.want {
                t.Errorf("got ID %d, want %d", got.ID, tt.want)
            }
        })
    }
}
```

---

## Day 10 — Interview Prep & Polish
**Concept:** Context, middleware, panic/recover + common Go interview questions  
**Build:** Add request logging middleware and graceful shutdown

### What to build
```go
// Logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

// Graceful shutdown
ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
defer cancel()
server := &http.Server{Addr: ":8080", Handler: mux}
go server.ListenAndServe()
<-ctx.Done()
server.Shutdown(context.Background())
```

### Common Go Interview Questions to Practice
1. What's the difference between a slice and an array?
2. When do you use a pointer receiver vs a value receiver?
3. What is an interface in Go? How does implicit satisfaction work?
4. How do goroutines differ from OS threads?
5. What is a channel? What's the difference between buffered and unbuffered?
6. How do you handle errors in Go? What is error wrapping?
7. What does `defer` do? When is it useful?
8. What is `nil` in Go? Can an interface be nil?
9. What's the zero value of a map? Why does writing to a nil map panic?
10. How does Go's garbage collector work (high level)?

---

## Quick Reference Cheatsheet

```go
// Slice
s := []int{1, 2, 3}
s = append(s, 4)
for i, v := range s { ... }

// Map
m := map[string]int{"a": 1}
v, ok := m["a"]

// Struct + method
type Dog struct{ Name string }
func (d *Dog) Bark() string { return d.Name + " says woof" }

// Interface
type Animal interface { Bark() string }
var a Animal = &Dog{Name: "Rex"}

// Goroutine + channel
ch := make(chan int, 1)
go func() { ch <- 42 }()
val := <-ch

// Error
if err != nil { return fmt.Errorf("context: %w", err) }

// Defer
defer file.Close() // runs when enclosing function returns
```
