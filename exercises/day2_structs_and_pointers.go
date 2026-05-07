// Day 2 Exercise: Structs & Pointers
//
// Run this file with: go run exercises/day2_structs_and_pointers.go
// (from inside your go-tasks folder)
//
// Work through each TODO. Expected output is shown above each section.
// Uncomment the answers at the bottom when you're done to check your work.

package main

import (
	"fmt"
	"time"
)

// -----------------------------------------------------------------
// These types are defined at the package level (outside any function)
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
}

func main() {
	fmt.Println("=== Day 2: Structs & Pointers ===")
	fmt.Println()

	// -----------------------------------------------------------------
	// PART 1: Creating structs
	// -----------------------------------------------------------------
	// Expected output:
	//   {1 Buy milk Buy milk from the store false medium 2009-11-10 23:00:00 +0000 UTC}
	//   Title: Buy milk
	//   Priority: medium
	//   Done: false

	fmt.Println("--- Part 1: Creating structs ---")

	// TODO: Create a Task using named fields (recommended style):
	//   ID: 1, Title: "Buy milk", Description: "Buy milk from the store",
	//   Done: false, Priority: Medium, CreatedAt: time.Now()
	// t1 := Task{
	// 	ID:          1,
	// 	Title:       "Buy milk",
	// 	Description: "Buy milk from the store",
	// 	Done:        false,
	// 	Priority:    Medium,
	// 	CreatedAt:   time.Now(),
	// }

	// TODO: Print the whole struct, then print Title, Priority, and Done separately
	// fmt.Println(t1)
	// fmt.Println("Title:", t1.Title)
	// fmt.Println("Priority:", t1.Priority)
	// fmt.Println("Done:", t1.Done)

	fmt.Println()

	// -----------------------------------------------------------------
	// PART 2: Zero values
	// -----------------------------------------------------------------
	// In Go every field gets a zero value if you don't set it:
	//   int → 0, string → "", bool → false, time.Time → zero time
	//
	// Expected output:
	//   Zero Task: {0  false  0001-01-01 00:00:00 +0000 UTC}
	//   ID is zero: true

	fmt.Println("--- Part 2: Zero values ---")

	// TODO: Create an empty Task with no fields set
	// var t2 Task
	// fmt.Println("Zero Task:", t2)
	// fmt.Println("ID is zero:", t2.ID == 0)

	fmt.Println()

	// -----------------------------------------------------------------
	// PART 3: Modifying struct fields (value vs pointer)
	// -----------------------------------------------------------------
	// This is the most important part of today. Pay close attention.
	//
	// Expected output:
	//   Before markDoneByValue:  false
	//   After markDoneByValue:   false   ← the original is unchanged!
	//   Before markDoneByPointer: false
	//   After markDoneByPointer:  true   ← pointer mutated the original

	fmt.Println("--- Part 3: Value vs Pointer ---")

	t3 := Task{ID: 2, Title: "Learn Go", Priority: High}

	// TODO: Call markDoneByValue(t3) and print t3.Done before and after.
	//       Notice the original does NOT change.
	// fmt.Println("Before markDoneByValue: ", t3.Done)
	// markDoneByValue(t3)
	// fmt.Println("After markDoneByValue:  ", t3.Done)

	// TODO: Call markDoneByPointer(&t3) and print t3.Done before and after.
	//       Notice the original DOES change.
	// fmt.Println("Before markDoneByPointer:", t3.Done)
	// markDoneByPointer(&t3)
	// fmt.Println("After markDoneByPointer: ", t3.Done)

	_ = t3 // remove once you use t3 above

	fmt.Println()

	// -----------------------------------------------------------------
	// PART 4: Working with pointers directly
	// -----------------------------------------------------------------
	// Expected output:
	//   Pointer address: 0xc000...  (some memory address)
	//   Value via pointer: Learn Go
	//   After update: Walk the dog
	//   Go auto-dereferences: Walk the dog  ← no need for (*p).Title

	fmt.Println("--- Part 4: Pointer mechanics ---")

	t4 := Task{ID: 3, Title: "Learn Go"}

	// TODO: Create a pointer to t4 using &
	// p := &t4

	// TODO: Print the pointer itself (shows memory address)
	// fmt.Println("Pointer address:", p)

	// TODO: Print the Title by dereferencing the pointer (*p).Title
	// fmt.Println("Value via pointer:", (*p).Title)

	// TODO: Update the Title through the pointer, then print t4.Title
	// p.Title = "Walk the dog"  // Go lets you skip the * for struct fields
	// fmt.Println("After update:", t4.Title)

	// TODO: Show that p.Title and (*p).Title are identical
	// fmt.Println("Go auto-dereferences:", p.Title)

	_ = t4 // remove once you use t4 above

	fmt.Println()

	// -----------------------------------------------------------------
	// PART 5: Slice of structs
	// -----------------------------------------------------------------
	// Expected output:
	//   All tasks:
	//     [1] Write tests (high) - done: false
	//     [2] Deploy app (medium) - done: false
	//     [3] Fix bug (low) - done: false
	//   After completing task 2:
	//     [2] Deploy app (medium) - done: true

	fmt.Println("--- Part 5: Slice of structs ---")

	tasks := []Task{
		{ID: 1, Title: "Write tests", Priority: High},
		{ID: 2, Title: "Deploy app", Priority: Medium},
		{ID: 3, Title: "Fix bug", Priority: Low},
	}

	// TODO: Print all tasks using a for-range loop and fmt.Printf
	//       Format: [ID] Title (priority) - done: Done
	// fmt.Println("All tasks:")
	// for _, t := range tasks {
	// 	fmt.Printf("  [%d] %s (%s) - done: %v\n", t.ID, t.Title, t.Priority, t.Done)
	// }

	// TODO: Mark the task with ID 2 as Done using its index (tasks[1].Done = true)
	//       then print just that task to confirm the change
	// tasks[1].Done = true
	// fmt.Println("After completing task 2:")
	// fmt.Printf("  [%d] %s (%s) - done: %v\n", tasks[1].ID, tasks[1].Title, tasks[1].Priority, tasks[1].Done)

	_ = tasks // remove once you use tasks above

	fmt.Println()

	// -----------------------------------------------------------------
	// PART 6: Challenge
	// -----------------------------------------------------------------
	// Write a function called updatePriority that takes a *Task and a
	// Priority, and sets the task's priority to the new value.
	// Then call it on one of your tasks and verify the change.
	//
	// Bonus: write a second version that takes a Task (value, not pointer)
	// and confirm the original is unchanged after calling it.

	fmt.Println("--- Part 6: Challenge ---")
	// Your code here...

	fmt.Println()
	fmt.Println("=== Done! Check the answers section below ===")
}

// -----------------------------------------------------------------
// These functions are used in Part 3.
// Notice: one takes Task (value copy), the other takes *Task (pointer).
// -----------------------------------------------------------------

// markDoneByValue receives a COPY of the task — changes don't affect the original
func markDoneByValue(t Task) {
	t.Done = true // modifies the local copy only
}

// markDoneByPointer receives the memory address — changes affect the original
func markDoneByPointer(t *Task) {
	t.Done = true // modifies the value at the address
}

// =============================================================================
// ANSWERS — uncomment to check your work
// =============================================================================

// Part 1:
// t1 := Task{
// 	ID:          1,
// 	Title:       "Buy milk",
// 	Description: "Buy milk from the store",
// 	Done:        false,
// 	Priority:    Medium,
// 	CreatedAt:   time.Now(),
// }
// fmt.Println(t1)
// fmt.Println("Title:", t1.Title)
// fmt.Println("Priority:", t1.Priority)
// fmt.Println("Done:", t1.Done)

// Part 2:
// var t2 Task
// fmt.Println("Zero Task:", t2)
// fmt.Println("ID is zero:", t2.ID == 0)

// Part 3:
// fmt.Println("Before markDoneByValue: ", t3.Done)
// markDoneByValue(t3)
// fmt.Println("After markDoneByValue:  ", t3.Done)
// fmt.Println("Before markDoneByPointer:", t3.Done)
// markDoneByPointer(&t3)
// fmt.Println("After markDoneByPointer: ", t3.Done)

// Part 4:
// p := &t4
// fmt.Println("Pointer address:", p)
// fmt.Println("Value via pointer:", (*p).Title)
// p.Title = "Walk the dog"
// fmt.Println("After update:", t4.Title)
// fmt.Println("Go auto-dereferences:", p.Title)

// Part 5:
// fmt.Println("All tasks:")
// for _, t := range tasks {
// 	fmt.Printf("  [%d] %s (%s) - done: %v\n", t.ID, t.Title, t.Priority, t.Done)
// }
// tasks[1].Done = true
// fmt.Println("After completing task 2:")
// fmt.Printf("  [%d] %s (%s) - done: %v\n", tasks[1].ID, tasks[1].Title, tasks[1].Priority, tasks[1].Done)

// Part 6 — challenge answer:
// func updatePriority(t *Task, p Priority) {
// 	t.Priority = p
// }
