// Day 3 Exercise: Methods
//
// Run this file with: go run exercises/day3_methods.go
// (from inside your go-tasks folder)
//
// Methods are functions attached to a type. Today you'll learn the difference
// between value receivers and pointer receivers, and when to use each.
//
// Work through each TODO. Expected output is shown above each section.
// Uncomment the answers at the bottom when you're done.

package main

/*
import (
	"fmt"
	"strings"
	"time"
)

// -----------------------------------------------------------------
// Types defined at the package level
// -----------------------------------------------------------------

type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
	Priority    Priority
	CreatedAt   time.Time
	Tags        []string
}

// -----------------------------------------------------------------
// PART 1: Your first method — value receiver
// -----------------------------------------------------------------
// A value receiver gets a COPY of the struct. Use it when:
//   - you only need to read fields (not mutate)
//   - the struct is small and copying is cheap
//
// Syntax: func (receiverName TypeName) MethodName() ReturnType { ... }

// TODO: Write a String() method on Task with a VALUE receiver.
//       It should return a string formatted like: "[1] Buy milk (high)"

var task Task = Task{ID: 1, Title: "Buy milk", Description: "Go buy milk", Done: false, Priority: Medium, CreatedAt: time.Now()}

func (t Task) String() string {
	return fmt.Sprintf("[%d] %s (%s)", t.ID, t.Title, t.Priority)
}

// -----------------------------------------------------------------
// PART 2: Pointer receiver methods
// -----------------------------------------------------------------
// A pointer receiver gets the memory address. Use it when:
//   - you need to mutate the struct
//   - the struct is large (avoids copying)
//   - consistency: if ANY method uses a pointer receiver, use pointer
//     receivers for ALL methods on that type (Go convention)

// TODO: Write a Complete() method on Task with a POINTER receiver.
//       It should set Done to true and record CreatedAt as time.Now().

func (t *Task) Complete() {
	t.Done = true;
}

// TODO: Write an AddTag() method on Task with a POINTER receiver.
//       It should append a tag string to the Tags slice.

func (t *Task) AddTag(tag string) {
	t.Tags = append(t.Tags, tag)
}


// TODO: Write a HasTag() method on Task with a VALUE receiver.
//       It should return true if the given tag exists in Tags.


func (t Task) HasTag(tag string) bool {
	for _, val := range t.Tags {
		if val ==  tag {
			return true
		}
	}
	return false
}

// -----------------------------------------------------------------
// PART 3: Methods on a non-struct type
// -----------------------------------------------------------------
// You can define methods on ANY named type, not just structs.
// Here we add a method to Priority (which is just a string underneath).

// TODO: Write an IsValid() method on Priority with a VALUE receiver.
//       It should return true if the priority is "low", "medium", or "high".
//

func (p Priority) IsValid() bool {
	if p == "low" || p == "medium" || p == "high"{
		return true
	}
	return false
}

// TODO: Write a Label() method on Priority that returns a display-friendly
//       string: "🔴 High", "🟡 Medium", "🟢 Low", or "Unknown" for anything else.
//

func (p Priority) Label() string {
	switch p{
		case "low":
			return "🟢 Low"
		case "medium":
			return "🟡 Medium"
		case "hight":
			return "🔴 High"
		default:
			return "unknown"
	}
}

// -----------------------------------------------------------------
// PART 4: MemoryStore — methods on a struct with a slice field
// -----------------------------------------------------------------
// This is a preview of what you'll build in store/memory.go today.

type MemoryStore struct {
	tasks  []Task
	nextID int
}

// TODO: Write a NewMemoryStore() function (not a method) that returns a
//       *MemoryStore with nextID set to 1.
//

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{nextID: 1}
}


// TODO: Write an Add() method on *MemoryStore.
//       It should set the task's ID to nextID, increment nextID,
//       append the task to the tasks slice, and return the task.

func (m *MemoryStore) Add(t *Task) Task {
	t.ID = m.nextID
	m.nextID++
	m.tasks = append(m.tasks, *t)
	return *t
}

// TODO: Write a GetAll() method on *MemoryStore that returns []Task.
//

func (m *MemoryStore) GetAll() []Task {
	return m.tasks
}

// TODO: Write a GetByID() method on *MemoryStore.
//       It should return (*Task, bool) — a pointer to the matching task
//       and true if found, or nil and false if not.

func (m *MemoryStore) GetById(id int) (*Task, bool) {
	for _, val := range m.tasks {
		if(val.ID == id){
			return &val, true
		}
	}
	return nil, false
}

// =============================================================================

func (m *MemoryStore) RemoveById(id int) {
	tasks := []Task{}
	for _, val := range m.tasks { if val.ID != id { tasks = append(tasks, val)}}
	m.tasks = tasks
}

func (m MemoryStore) Count() int {
	return len(m.tasks)
}

func main() {

	fmt.Println("=== Day 3: Methods ===")
	fmt.Println()

	// -----------------------------------------------------------------
	// Part 1: String() — value receiver
	// -----------------------------------------------------------------
	// Expected output:
	//   [1] Learn Go (high)

	fmt.Println("--- Part 1: String() method ---")

	t1 := Task{ID: 1, Title: "Learn Go", Priority: High}

	// TODO: Call t1.String() and print the result
	// fmt.Println(task.String())

	_ = t1
	fmt.Println()

	// -----------------------------------------------------------------
	// Part 2: Pointer receiver methods
	// -----------------------------------------------------------------
	// Expected output:
	//   Done before: false
	//   Done after Complete(): true
	//   Tags before: []
	//   Tags after AddTag: [backend urgent]
	//   HasTag "urgent": true
	//   HasTag "frontend": false

	fmt.Println("--- Part 2: Pointer receiver methods ---")

	t2 := Task{ID: 2, Title: "Deploy app", Priority: Medium}

	// TODO: Print t2.Done, call t2.Complete(), print t2.Done again
	fmt.Println("Done before: ", t2.Done)
	t2.Complete()
	fmt.Println("Done after Complete(): ", t2.Done)

	// TODO: Print t2.Tags, add tags "backend" and "urgent", print t2.Tags again
	fmt.Println("Tags before: ", t2.Tags)
	t2.AddTag("backend")
	t2.AddTag("urgent")
	fmt.Println("Tags after AddTags: ", t2.Tags)


	// TODO: Call t2.HasTag("urgent") and t2.HasTag("frontend"), print both
	t2.HasTag("urgent")

	_ = t2
	fmt.Printf("HasTag %s: %v\n", "urgent" ,t2.HasTag("urgent"))
	fmt.Printf("HasTag %s: %v\n", "backend" ,t2.HasTag("backend"))

	// -----------------------------------------------------------------
	// Part 3: Methods on Priority type
	// -----------------------------------------------------------------
	// Expected output:
	//   high is valid: true
	//   urgent is valid: false
	//   High label: 🔴 High
	//   Medium label: 🟡 Medium


	fmt.Println("--- Part 3: Methods on Priority ---")

	// TODO: Call IsValid() on High and on Priority("urgent"), print results
	var highIsValid bool= Priority(High).IsValid()
	var urgentIsValid bool = Priority("urgent").IsValid()
	fmt.Println("high is valid: ", highIsValid)
	fmt.Println("urgent is valid: ", urgentIsValid)


	// TODO: Call Label() on High and Medium, print results
	fmt.Println("Medium label: ", Priority(Medium).Label())
	fmt.Println("High label: ", Priority(High).Label())

	fmt.Println()

	// -----------------------------------------------------------------
	// Part 4: MemoryStore
	// -----------------------------------------------------------------
	// Expected output:
	//   Added: [1] Write tests (high)
	//   Added: [2] Fix bug (low)
	//   All tasks: 2
	//     [1] Write tests (high) done: false
	//     [2] Fix bug (low) done: false
	//   Found by ID 1: Write tests
	//   ID 99 found: false

	fmt.Println("--- Part 4: MemoryStore ---")

	// TODO: Create a new store using NewMemoryStore()
	var m MemoryStore = *NewMemoryStore()
	// TODO: Add two tasks and print each using String()
	m.Add(&Task{Title: "Write tests", Priority: High, Done: false})
	m.Add(&Task{Title: "Fix bug", Priority: Low, Done: false})
	// TODO: Call GetAll(), print the count and each task
	var tasks []Task = m.GetAll();

	for _, val := range tasks {
		fmt.Println(val)
	}

	// TODO: Call GetByID(1) and GetByID(99), print results
	_, exists1 := m.GetById(1)
	_, exists2 := m.GetById(2)

	fmt.Println("Foudn by ID 1: ", exists1 )
	fmt.Println("ID 99 found: ", exists2)

	fmt.Println()

	// -----------------------------------------------------------------
	// Part 5: Challenge
	// -----------------------------------------------------------------
	// Add a Delete(id int) bool method to MemoryStore.
	// It should remove the task with the matching ID from the slice and
	// return true, or return false if no task with that ID exists.
	//
	// Hint: build a new slice, skipping the task you want to remove.
	//   var updated []Task
	//   for _, t := range s.tasks { if t.ID != id { updated = append(...) } }
	//
	// Then add a Count() int method that returns len(s.tasks).
	// Verify: add 3 tasks, delete one, confirm Count() is 2.

	fmt.Println("--- Part 5: Challenge ---")
	fmt.Println(m.Count())
	m.RemoveById(1)
	fmt.Println(m.Count())

	fmt.Println()
	fmt.Println("=== Done! Check the answers below ===")
}

// =============================================================================
// ANSWERS — uncomment to check your work
// =============================================================================

// Part 1:
// func (t Task) String() string {
// 	return fmt.Sprintf("[%d] %s (%s)", t.ID, t.Title, t.Priority)
// }
// fmt.Println(t1.String())

// Part 2:
// func (t *Task) Complete() {
// 	t.Done = true
// 	t.CreatedAt = time.Now()
// }
// func (t *Task) AddTag(tag string) {
// 	t.Tags = append(t.Tags, tag)
// }
// func (t Task) HasTag(tag string) bool {
// 	for _, tg := range t.Tags {
// 		if tg == tag { return true }
// 	}
// 	return false
// }

// Part 3:
// func (p Priority) IsValid() bool {
// 	return p == Low || p == Medium || p == High
// }
// func (p Priority) Label() string {
// 	switch p {
// 	case High:   return "🔴 High"
// 	case Medium: return "🟡 Medium"
// 	case Low:    return "🟢 Low"
// 	default:     return "Unknown"
// 	}
// }

// Part 4:
// func NewMemoryStore() *MemoryStore { return &MemoryStore{nextID: 1} }
// func (s *MemoryStore) Add(t Task) Task {
// 	t.ID = s.nextID; s.nextID++; s.tasks = append(s.tasks, t); return t
// }
// func (s *MemoryStore) GetAll() []Task { return s.tasks }
// func (s *MemoryStore) GetByID(id int) (*Task, bool) {
// 	for i := range s.tasks {
// 		if s.tasks[i].ID == id { return &s.tasks[i], true }
// 	}
// 	return nil, false
// }

// Part 5 — challenge answer:
// func (s *MemoryStore) Delete(id int) bool {
// 	var updated []Task
// 	found := false
// 	for _, t := range s.tasks {
// 		if t.ID == id { found = true; continue }
// 		updated = append(updated, t)
// 	}
// 	s.tasks = updated
// 	return found
// }
// func (s *MemoryStore) Count() int { return len(s.tasks) }

// Suppress unused import warning if Parts 1/2 are still commented out
var _ = strings.Contains
var _ = time.Now
*/