// Day 4 Exercise: Interfaces
//
// Run this file with: go run exercises/day4_interfaces.go
// (from inside your go-tasks folder)
//
// Interfaces are Go's most powerful feature for writing flexible, testable code.
// A type satisfies an interface simply by having the right methods — no
// "implements" keyword needed. This is called implicit satisfaction.
//
// Work through each TODO. Expected output is shown above each section.
// Uncomment the answers at the bottom when you're done.

package main

import (
	"fmt"
	"strings"
)

// =============================================================================
// PART 1: Defining and satisfying an interface
// =============================================================================
// An interface defines a set of method signatures.
// Any type that has those methods automatically satisfies the interface.

// Describable is a simple interface — anything that can describe itself.
type Describable interface {
	Describe() string
}

type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

type Task struct {
	ID       int
	Title    string
	Done     bool
	Priority Priority
	Tags     []string
}

type Project struct {
	Name      string
	TaskCount int
}

// TODO: Write a Describe() method on Task (value receiver) that returns:
//       "Task 1: Buy milk [high] - pending"   (if Done is false)
//       "Task 1: Buy milk [high] - completed" (if Done is true)
//
func (t Task) Describe() string {
	var complete string = "Pending"
	if t.Done { complete = "Complete"}
	return fmt.Sprintf("Task %d: %s %v - %s", t.ID, t.Title, t.Priority, complete)
}


// TODO: Write a Describe() method on Project (value receiver) that returns:
//       "Project: GoTraining (5 tasks)"

func (p Project) Describe() string {
	return fmt.Sprintf("%s [%d]", p.Name, p.TaskCount)
}

// printDescription takes the INTERFACE type — it works with any Describable.
// This is the power of interfaces: one function, many types.
func printDescription(d Describable) {
	fmt.Println(d.Describe())
}

// =============================================================================
// PART 2: The TaskStore interface — repository pattern
// =============================================================================
// This is exactly what you'll use in your real project.
// Handlers talk to a TaskStore interface, not a concrete type.
// That means you can swap implementations (memory → database) without
// touching any handler code.

type TaskStore interface {
	Add(task Task) Task
	GetAll() []Task
	GetByID(id int) (*Task, bool)
	Delete(id int) bool
	Count() int
}

// MemoryStore is a concrete type that will satisfy TaskStore.
type MemoryStore struct {
	tasks  []Task
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{nextID: 1}
}

// TODO: Implement all 5 methods on *MemoryStore so it satisfies TaskStore.

func (m *MemoryStore) Add(task Task) Task {
	newTasks := m.tasks[:]
	newTasks = append(newTasks, task)
	m.tasks = newTasks
	return task
}

// GetAll — return the slice

func (m *MemoryStore) GetAll() []Task {
	return m.tasks
}

// GetByID — loop with index, return &s.tasks[i] and true, or nil and false
func (m *MemoryStore) GetByID(id int) (*Task, bool) {
	for i := range m.tasks {
		if m.tasks[i].ID == id {
			return &m.tasks[i], true
		}
	}
	return nil, false
}

// Delete — rebuild the slice skipping the matching ID, return found bool

func (m *MemoryStore) Delete(id int) bool {
	newTasks := []Task{}
	lenBefore := len(m.tasks)
	for _, val := range m.tasks {
		if val.ID != id {
			newTasks = append(newTasks, val)
		}
	}
	m.tasks = newTasks
	return len(m.tasks) < lenBefore
}

// Count — return len(s.tasks)

func (m *MemoryStore) Count() int {
	return len(m.tasks)
}

// =============================================================================
// PART 3: Interface as a function parameter
// =============================================================================
// printSummary takes a TaskStore interface — it doesn't care whether it's
// a MemoryStore, a PostgresStore, or a mock. As long as it has the methods.

func printSummary(store TaskStore) {
	tasks := store.GetAll()
	fmt.Printf("Store has %d task(s):\n", store.Count())
	for _, t := range tasks {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("  %s [%d] %s (%s)\n", status, t.ID, t.Title, t.Priority)
	}
}

// =============================================================================
// PART 4: Interface compliance check (compile-time guarantee)
// =============================================================================
// This pattern forces a compile error if MemoryStore ever stops satisfying
// TaskStore — useful as your codebase grows.
//
// TODO: Uncomment this line. If your MemoryStore methods are all implemented
//       correctly, it will compile silently. If any method is missing, you'll
//       get a clear error pointing to exactly what's missing.
//
var _ TaskStore = (*MemoryStore)(nil)

// =============================================================================
// PART 5: The empty interface and type assertions
// =============================================================================
// interface{} (or 'any' in modern Go) accepts any type at all.
// To get the concrete value back out, you use a type assertion.

func describe(i interface{}) {
	// TODO: Use a type switch to print different messages based on the type:
	//   int    → "Got an int: <value>"
	//   string → "Got a string: <value>"
	//   Task   → "Got a Task: <title>"
	//   default → "Got something else"

	switch v := i.(type) {
	case int: 
		fmt.Println("Got an int: ", v)
	case string:
		fmt.Println("Got a string: ", v)
	case Task:
		fmt.Println("Got a Task ", v.Title)
	default:
		fmt.Println("Got something else")
	}
}

// =============================================================================

func main() {
	fmt.Println("=== Day 4: Interfaces ===")
	fmt.Println()

	// -----------------------------------------------------------------------
	// Part 1: Implicit interface satisfaction
	// -----------------------------------------------------------------------
	// Expected output:
	//   Task 1: Learn Go [high] - pending
	//   Task 2: Deploy app [medium] - completed
	//   Project: GoTraining (12 tasks)

	fmt.Println("--- Part 1: Describable interface ---")

	t1 := Task{ID: 1, Title: "Learn Go", Priority: High}
	t2 := Task{ID: 2, Title: "Deploy app", Priority: Medium, Done: true}
	p1 := Project{Name: "GoTraining", TaskCount: 12}

	// TODO: Call printDescription() with t1, t2, and p1.
	//       Notice one function handles all three types.
	printDescription(t1)
	printDescription(t2)
	printDescription(p1)

	_ = t1
	_ = t2
	_ = p1
	fmt.Println()

	// -----------------------------------------------------------------------
	// Part 2 & 3: TaskStore interface + printSummary
	// -----------------------------------------------------------------------
	// Expected output:
	//   Store has 3 task(s):
	//     [ ] [1] Write tests (high)
	//     [ ] [2] Fix bug (low)
	//     [x] [3] Deploy app (medium)
	//   After delete — store has 2 task(s):
	//     [ ] [1] Write tests (high)
	//     [x] [3] Deploy app (medium)

	fmt.Println("--- Part 2 & 3: TaskStore interface ---")

	// TODO: Create a store and assign it to a TaskStore variable (not *MemoryStore).
	//       This proves MemoryStore satisfies the interface.
	var newStore TaskStore = NewMemoryStore()
	// TODO: Add three tasks
	newStore.Add(Task{ID: 0, Title: "Write tests", Priority: High })
	newStore.Add(Task{ID: 1, Title: "Fix bug", Priority: Low})
	newStore.Add(Task{ID: 2, Title: "Deploy app", Priority: Medium})	

	// TODO: Mark the third task done via GetByID, then call printSummary
	if found, ok := newStore.GetByID(2); ok { found.Done = true }
	printSummary(newStore)
	
	// TODO: Delete task ID 2, print separator, call printSummary again
	newStore.Delete(2)
	fmt.Print("After delete — ")
	printSummary(newStore)
	fmt.Println()

	// -----------------------------------------------------------------------
	// Part 4: Compile-time interface check
	// -----------------------------------------------------------------------
	// Expected output: (nothing — it's a compile-time check only)

	fmt.Println("--- Part 4: Compile-time check ---")
	// Uncomment the var _ line above main() and run — if it compiles, you're good.
	fmt.Println("If this compiled, MemoryStore satisfies TaskStore ✓")
	fmt.Println()

	// -----------------------------------------------------------------------
	// Part 5: Empty interface and type assertions
	// -----------------------------------------------------------------------
	// Expected output:
	//   Got an int: 42
	//   Got a string: hello
	//   Got a Task: Learn Go
	//   Got something else

	fmt.Println("--- Part 5: Empty interface & type switch ---")

	// TODO: Fill in the describe() function above, then call it with these values
	describe(42)
	describe("hello")
	describe(Task{ID:0, Title: "Learn Go"})
	describe(false)

	fmt.Println()

	// -----------------------------------------------------------------------
	// Part 6: Challenge
	// -----------------------------------------------------------------------
	// 1. Create a second store type called FilterStore that wraps a TaskStore
	//    and adds a GetByTag(tag string) []Task method.
	//
	// 2. Define a new interface called TaggableStore that embeds TaskStore
	//    and adds GetByTag.
	//
	// 3. Implement GetByTag on FilterStore by iterating over store.GetAll()
	//    and returning tasks whose Tags slice contains the given tag.
	//
	// Hint for embedding interfaces:
	//   type TaggableStore interface {
	//       TaskStore                    // embed — inherits all TaskStore methods
	//       GetByTag(tag string) []Task  // adds one more
	//   }

	fmt.Println("--- Part 6: Challenge ---")
	// Your code here...

	fmt.Println()
	fmt.Println("=== Done! Check the answers below ===")
}

// =============================================================================
// ANSWERS — uncomment to check your work
// =============================================================================

// Part 1:
// func (t Task) Describe() string {
// 	status := "pending"
// 	if t.Done { status = "completed" }
// 	return fmt.Sprintf("Task %d: %s [%s] - %s", t.ID, t.Title, t.Priority, status)
// }
// func (p Project) Describe() string {
// 	return fmt.Sprintf("Project: %s (%d tasks)", p.Name, p.TaskCount)
// }
// printDescription(t1)
// printDescription(t2)
// printDescription(p1)

// Part 2 (MemoryStore methods): see method stubs in Part 2 section above

// Part 3:
// var store TaskStore = NewMemoryStore()
// store.Add(Task{Title: "Write tests", Priority: High})
// store.Add(Task{Title: "Fix bug", Priority: Low})
// t3 := store.Add(Task{Title: "Deploy app", Priority: Medium})
// if found, ok := store.GetByID(t3.ID); ok { found.Done = true }
// printSummary(store)
// store.Delete(2)
// fmt.Print("After delete — ")
// printSummary(store)

// Part 5:
// switch v := i.(type) {
// case int:    fmt.Println("Got an int:", v)
// case string: fmt.Println("Got a string:", v)
// case Task:   fmt.Println("Got a Task:", v.Title)
// default:     fmt.Println("Got something else")
// }

// Part 6 — challenge answer:
// type FilterStore struct{ store TaskStore }
// type TaggableStore interface {
// 	TaskStore
// 	GetByTag(tag string) []Task
// }
// func (f *FilterStore) Add(t Task) Task           { return f.store.Add(t) }
// func (f *FilterStore) GetAll() []Task            { return f.store.GetAll() }
// func (f *FilterStore) GetByID(id int) (*Task, bool) { return f.store.GetByID(id) }
// func (f *FilterStore) Delete(id int) bool        { return f.store.Delete(id) }
// func (f *FilterStore) Count() int                { return f.store.Count() }
// func (f *FilterStore) GetByTag(tag string) []Task {
// 	var result []Task
// 	for _, t := range f.store.GetAll() {
// 		for _, tg := range t.Tags {
// 			if tg == tag { result = append(result, t); break }
// 		}
// 	}
// 	return result
// }

// Suppress unused import warning
var _ = strings.Contains
