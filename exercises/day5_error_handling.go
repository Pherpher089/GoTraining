// Day 5 Exercise: Error Handling
//
// Run this file with: go run exercises/day5_error_handling.go
// (from inside your go-tasks folder)
//
// Go handles errors as values, not exceptions. There is no try/catch.
// Instead, functions return an error as their last return value, and
// callers check it immediately. This makes error paths explicit and
// impossible to accidentally ignore.
//
// Work through each TODO. Expected output is shown above each section.
// Uncomment the answers at the bottom when you're done.

package main

import (
	"errors"
	"fmt"
)

// =============================================================================
// PART 1: The error interface
// =============================================================================
// error is just a built-in interface:
//   type error interface {
//       Error() string
//   }
//
// Any type with an Error() string method satisfies it.
// The simplest errors are created with errors.New() or fmt.Errorf().

// TODO: Write a function called divide(a, b float64) (float64, error)
//       If b is 0, return 0 and an error created with errors.New("cannot divide by zero")
//       Otherwise return a/b and nil
//
// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, errors.New("cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// =============================================================================
// PART 2: Sentinel errors
// =============================================================================
// A sentinel error is a package-level error variable used for comparison.
// You've seen this pattern: io.EOF, sql.ErrNoRows, etc.
// Callers can check: if err == ErrNotFound { ... }
// Better yet, use errors.Is() which handles wrapping (more on that in Part 4).

// TODO: Declare two sentinel errors at the package level:
//       ErrNotFound — "task not found"
//       ErrEmptyTitle — "task title cannot be empty"
//
// var ErrNotFound = errors.New("task not found")
// var ErrEmptyTitle = errors.New("task title cannot be empty")

// =============================================================================
// PART 3: Custom error types
// =============================================================================
// Sometimes you need to attach extra data to an error (like which ID wasn't found).
// Do this by defining a struct that implements the error interface.

// TODO: Define a NotFoundError struct with an ID int field.
//       Implement Error() string returning: "task with ID 42 not found"
//
// type NotFoundError struct {
// 	ID int
// }
// func (e *NotFoundError) Error() string {
// 	return fmt.Sprintf("task with ID %d not found", e.ID)
// }

// TODO: Define a ValidationError struct with Field string and Message string fields.
//       Implement Error() string returning: "validation error on title: cannot be empty"
//
// type ValidationError struct {
// 	Field   string
// 	Message string
// }
// func (e *ValidationError) Error() string {
// 	return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
// }

// =============================================================================
// PART 4: Error wrapping with fmt.Errorf and %w
// =============================================================================
// Wrapping adds context to an error while preserving the original underneath.
// Use %w (not %v) to wrap so callers can unwrap with errors.Is / errors.As.
//
// errors.Is(err, target) — checks if target appears anywhere in the error chain
// errors.As(err, &target) — extracts a specific error type from the chain

// TODO: Write a function getTask(id int) (*string, error) that:
//       - returns an error if id <= 0: fmt.Errorf("getTask: %w", ErrNotFound)
//       - returns a pointer to "task data" and nil otherwise
//
// func getTask(id int) (*string, error) {
// 	if id <= 0 {
// 		return nil, fmt.Errorf("getTask: %w", ErrNotFound)
// 	}
// 	result := "task data"
// 	return &result, nil
// }

// =============================================================================
// PART 5: A store with proper error handling
// =============================================================================
// This mirrors what your real store/memory.go should look like after Day 5.

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
}

type MemoryStore struct {
	tasks  []Task
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{nextID: 1}
}

// TODO: Write Add(task Task) (Task, error)
//       Return a ValidationError if task.Title is empty.
//       Otherwise assign ID, increment nextID, append, return task and nil.
//
// func (s *MemoryStore) Add(task Task) (Task, error) {
// 	if task.Title == "" {
// 		return Task{}, &ValidationError{Field: "title", Message: "cannot be empty"}
// 	}
// 	task.ID = s.nextID
// 	s.nextID++
// 	s.tasks = append(s.tasks, task)
// 	return task, nil
// }

// TODO: Write GetByID(id int) (*Task, error)
//       Return &NotFoundError{ID: id} if not found.
//       Otherwise return a pointer to the task and nil.
//       Remember: range by index so the pointer stays valid.
//
// func (s *MemoryStore) GetByID(id int) (*Task, error) {
// 	for i := range s.tasks {
// 		if s.tasks[i].ID == id {
// 			return &s.tasks[i], nil
// 		}
// 	}
// 	return nil, &NotFoundError{ID: id}
// }

// TODO: Write Delete(id int) error
//       Return &NotFoundError{ID: id} if no task with that ID exists.
//       Otherwise rebuild the slice without that task and return nil.
//
// func (s *MemoryStore) Delete(id int) error {
// 	var updated []Task
// 	found := false
// 	for _, t := range s.tasks {
// 		if t.ID == id {
// 		    found = true
// 		    continue
// 		}
// 		updated = append(updated, t)
// 	}
// 	if !found {
// 		return &NotFoundError{ID: id}
// 	}
// 	s.tasks = updated
// 	return nil
// }

// =============================================================================

func main() {
	fmt.Println("=== Day 5: Error Handling ===")
	fmt.Println()

	// -----------------------------------------------------------------------
	// Part 1: Basic error returns
	// -----------------------------------------------------------------------
	// Expected output:
	//   10 / 2 = 5.00
	//   10 / 0: cannot divide by zero

	fmt.Println("--- Part 1: Basic error returns ---")

	// TODO: Call divide(10, 2), print the result
	//       Call divide(10, 0), print the error
	//
	// result, err := divide(10, 2)
	// if err != nil {
	// 	fmt.Println("10 / 2 error:", err)
	// } else {
	// 	fmt.Printf("10 / 2 = %.2f\n", result)
	// }
	//
	// _, err = divide(10, 0)
	// if err != nil {
	// 	fmt.Println("10 / 0:", err)
	// }

	fmt.Println()

	// -----------------------------------------------------------------------
	// Part 2: Sentinel errors with errors.Is
	// -----------------------------------------------------------------------
	// Expected output:
	//   Error: task not found
	//   Is ErrNotFound: true
	//   Is ErrEmptyTitle: false

	fmt.Println("--- Part 2: Sentinel errors ---")

	// TODO: Create an err variable set to ErrNotFound
	//       Use errors.Is to check if it matches ErrNotFound and ErrEmptyTitle
	//
	// err := ErrNotFound
	// fmt.Println("Error:", err)
	// fmt.Println("Is ErrNotFound:", errors.Is(err, ErrNotFound))
	// fmt.Println("Is ErrEmptyTitle:", errors.Is(err, ErrEmptyTitle))

	fmt.Println()

	// -----------------------------------------------------------------------
	// Part 3: Custom error types with errors.As
	// -----------------------------------------------------------------------
	// Expected output:
	//   Error: task with ID 7 not found
	//   Is a NotFoundError: true
	//   The missing ID was: 7
	//   Error: validation error on title: cannot be empty
	//   Is a ValidationError: true
	//   Bad field: title

	fmt.Println("--- Part 3: Custom error types ---")

	// TODO: Create a *NotFoundError for ID 7, print it
	//       Use errors.As to extract it and print the ID field
	//
	// var nfe *NotFoundError
	// err1 := &NotFoundError{ID: 7}
	// fmt.Println("Error:", err1)
	// fmt.Println("Is a NotFoundError:", errors.As(err1, &nfe))
	// fmt.Println("The missing ID was:", nfe.ID)

	// TODO: Create a *ValidationError, print it
	//       Use errors.As to extract it and print the Field
	//
	// var ve *ValidationError
	// err2 := &ValidationError{Field: "title", Message: "cannot be empty"}
	// fmt.Println("Error:", err2)
	// fmt.Println("Is a ValidationError:", errors.As(err2, &ve))
	// fmt.Println("Bad field:", ve.Field)

	fmt.Println()

	// -----------------------------------------------------------------------
	// Part 4: Wrapped errors and error chains
	// -----------------------------------------------------------------------
	// Expected output:
	//   Error: getTask: task not found
	//   errors.Is still finds ErrNotFound through the wrap: true

	fmt.Println("--- Part 4: Error wrapping ---")

	// TODO: Call getTask(-1) and print the error
	//       Then use errors.Is to confirm ErrNotFound is in the chain
	//       despite the wrapping
	//
	// _, err := getTask(-1)
	// fmt.Println("Error:", err)
	// fmt.Println("errors.Is still finds ErrNotFound through the wrap:", errors.Is(err, ErrNotFound))

	fmt.Println()

	// -----------------------------------------------------------------------
	// Part 5: Store with error handling
	// -----------------------------------------------------------------------
	// Expected output:
	//   Added: [1] Learn Go
	//   Add empty title error: validation error on title: cannot be empty
	//   Is ValidationError: true
	//   Found task: Learn Go
	//   GetByID(99) error: task with ID 99 not found
	//   Is NotFoundError: true, ID: 99
	//   Deleted task 1: <nil>
	//   Delete(1) again: task with ID 1 not found

	fmt.Println("--- Part 5: Store with error handling ---")

	store := NewMemoryStore()

	// TODO: Add a valid task and print it
	// t, err := store.Add(Task{Title: "Learn Go", Priority: High})
	// if err != nil {
	// 	fmt.Println("Add error:", err)
	// } else {
	// 	fmt.Printf("Added: [%d] %s\n", t.ID, t.Title)
	// }

	// TODO: Try adding a task with an empty title, print the error
	//       Use errors.As to confirm it's a ValidationError
	// _, err = store.Add(Task{})
	// fmt.Println("Add empty title error:", err)
	// var ve *ValidationError
	// fmt.Println("Is ValidationError:", errors.As(err, &ve))

	// TODO: GetByID(1) — print the task title
	// found, err := store.GetByID(1)
	// if err != nil {
	// 	fmt.Println("GetByID error:", err)
	// } else {
	// 	fmt.Println("Found task:", found.Title)
	// }

	// TODO: GetByID(99) — print the error, use errors.As to get the NotFoundError ID
	// _, err = store.GetByID(99)
	// fmt.Println("GetByID(99) error:", err)
	// var nfe *NotFoundError
	// if errors.As(err, &nfe) {
	// 	fmt.Println("Is NotFoundError: true, ID:", nfe.ID)
	// }

	// TODO: Delete(1) — print the error (should be nil)
	//       Delete(1) again — print the error (should be NotFoundError)
	// err = store.Delete(1)
	// fmt.Println("Deleted task 1:", err)
	// err = store.Delete(1)
	// fmt.Println("Delete(1) again:", err)

	_ = store
	fmt.Println()

	// -----------------------------------------------------------------------
	// Part 6: Challenge
	// -----------------------------------------------------------------------
	// Add an Update(id int, updated Task) (Task, error) method to MemoryStore.
	// It should:
	//   - return &NotFoundError{ID: id} if the task doesn't exist
	//   - return &ValidationError if updated.Title is empty
	//   - otherwise replace the task in the slice and return the updated task
	//
	// Then write a helper handleStoreError(err error) string that returns:
	//   "not found"   if errors.As finds a *NotFoundError
	//   "bad request" if errors.As finds a *ValidationError
	//   "internal"    for anything else
	// This mirrors how an HTTP handler would map errors to status codes.

	fmt.Println("--- Part 6: Challenge ---")
	// Your code here...

	fmt.Println()
	fmt.Println("=== Done! Check the answers below ===")
}

// =============================================================================
// ANSWERS — uncomment to check your work
// =============================================================================

// Part 1:
// func divide(a, b float64) (float64, error) {
// 	if b == 0 { return 0, errors.New("cannot divide by zero") }
// 	return a / b, nil
// }

// Part 2:
// var ErrNotFound = errors.New("task not found")
// var ErrEmptyTitle = errors.New("task title cannot be empty")

// Part 3:
// type NotFoundError struct{ ID int }
// func (e *NotFoundError) Error() string { return fmt.Sprintf("task with ID %d not found", e.ID) }
// type ValidationError struct{ Field, Message string }
// func (e *ValidationError) Error() string {
// 	return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
// }

// Part 4:
// func getTask(id int) (*string, error) {
// 	if id <= 0 { return nil, fmt.Errorf("getTask: %w", ErrNotFound) }
// 	result := "task data"
// 	return &result, nil
// }

// Part 5 store methods: see stubs in Part 5 section above

// Part 6 — challenge answer:
// func (s *MemoryStore) Update(id int, updated Task) (Task, error) {
// 	if updated.Title == "" {
// 		return Task{}, &ValidationError{Field: "title", Message: "cannot be empty"}
// 	}
// 	for i := range s.tasks {
// 		if s.tasks[i].ID == id {
// 			updated.ID = id
// 			s.tasks[i] = updated
// 			return s.tasks[i], nil
// 		}
// 	}
// 	return Task{}, &NotFoundError{ID: id}
// }
//
// func handleStoreError(err error) string {
// 	var nfe *NotFoundError
// 	var ve *ValidationError
// 	switch {
// 	case errors.As(err, &nfe): return "not found"
// 	case errors.As(err, &ve):  return "bad request"
// 	default:                   return "internal"
// 	}
// }

// Suppress unused import
var _ = errors.New
