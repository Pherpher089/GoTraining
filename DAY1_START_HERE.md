# Day 1 — Start Here

Follow these steps yourself in your terminal. Each command teaches you something real about how Go projects are structured.

---

## Step 1: Create the project folder

```bash
cd path/to/GoTraining   # navigate to your GoTraining folder
mkdir go-tasks
cd go-tasks
```

---

## Step 2: Initialize the Go module

```bash
go mod init github.com/chris/go-tasks
```

This creates a `go.mod` file — Go's equivalent of `package.json`. The module path (`github.com/chris/go-tasks`) is how other files in your project import each other. It doesn't have to be a real GitHub URL, but it's conventional to use one.

Open `go.mod` and take a look. You'll see the module name and the Go version.

---

## Step 3: Create your folder structure

```bash
mkdir models store handlers exercises
```

Your project will look like this by the end:

```
go-tasks/
  go.mod
  main.go
  models/
    task.go          ← Task struct (Day 2)
  store/
    store.go         ← TaskStore interface (Day 4)
    memory.go        ← In-memory implementation (Day 3)
  handlers/
    task_handler.go  ← HTTP handlers (Day 6)
  exercises/
    day1_slices.go   ← today's exercise (already written for you)
```

---

## Step 4: Write your first Go file

Create `main.go` in the `go-tasks/` folder:

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

**Things to notice:**
- `package main` — every executable Go file lives in `package main`
- `import` — Go's import block; unused imports are a compile error
- `http.HandleFunc` — registers a route handler
- `log.Fatal` — prints and exits if the server fails to start
- The anonymous function `func(w http.ResponseWriter, r *http.Request)` is Go's handler signature — you'll write this a lot

---

## Step 5: Run it

```bash
go run main.go
```

Visit `http://localhost:8080` in your browser. You should see: **go-tasks API is running!**

Press `Ctrl+C` to stop.

---

## Step 6: Do the Day 1 exercise

Copy the exercise file from the GoTraining root into your project's exercises folder:

```bash
cp ../exercises/day1_slices.go exercises/day1_slices.go
```

Then work through it:

```bash
go run exercises/day1_slices.go
```

The file has TODOs for you to fill in. Read the expected output in the comments, write the code, run it, check it matches. Uncomment the answer section at the bottom when you're done.

---

## Useful Go CLI commands to know

| Command | What it does |
|---|---|
| `go run main.go` | Compile + run in one step |
| `go build` | Compile to a binary |
| `go test ./...` | Run all tests in the project |
| `go fmt ./...` | Auto-format all code |
| `go vet ./...` | Catch common mistakes |
| `go mod tidy` | Clean up unused dependencies |
| `go get <pkg>` | Add a dependency |

---

## What's next (Day 2)

After finishing the exercise, read the **Day 2** section in `TRAINING_PLAN.md`. You'll define the `Task` struct and start learning about structs and pointers. Come back to Cowork and I'll help you build it!
