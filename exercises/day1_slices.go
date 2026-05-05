// Day 1 Exercise: Arrays & Slices
//
// Run this file with: go run exercises/day1_slices.go
// (from inside the go-tasks folder after you've initialized the module)
//
// Work through each TODO. The expected output is shown in comments above each section.
// Uncomment the answer blocks at the bottom to check your work.

package main

import "fmt"

func main() {
	fmt.Println("=== Day 1: Arrays & Slices ===")
	fmt.Println()

	// -----------------------------------------------------------------
	// PART 1: Arrays (fixed size)
	// -----------------------------------------------------------------
	// Expected output:
	//   Array: [0 0 0 0 0]
	//   After setting index 2: [0 0 99 0 0]
	//   Length: 5

	fmt.Println("--- Part 1: Arrays ---")

	// TODO: Declare an array of 5 ints named `scores`
	// var scores [5]int

	// TODO: Set the third element (index 2) to 99
	// scores[2] = 99

	// TODO: Print the array, then print its length using len()
	// fmt.Println("Array:", scores)
	// fmt.Println("After setting index 2:", scores)
	// fmt.Println("Length:", len(scores))

	fmt.Println()

	// -----------------------------------------------------------------
	// PART 2: Slices (dynamic — what you'll use in real code)
	// -----------------------------------------------------------------
	// Expected output:
	//   Tasks: [Write tests Deploy app Review PR]
	//   Length: 3, Capacity: 4 (or similar — capacity may vary)
	//   First two: [Write tests Deploy app]

	fmt.Println("--- Part 2: Slices ---")

	// TODO: Create a slice of strings with these three values:
	//   "Write tests", "Deploy app", "Review PR"
	// tasks := []string{"Write tests", "Deploy app", "Review PR"}

	// TODO: Print the slice
	// fmt.Println("Tasks:", tasks)

	// TODO: Print its length (len) and capacity (cap)
	// fmt.Printf("Length: %d, Capacity: %d\n", len(tasks), cap(tasks))

	// TODO: Print just the first two elements using a slice expression
	// fmt.Println("First two:", tasks[:2])

	fmt.Println()

	// -----------------------------------------------------------------
	// PART 3: append and range
	// -----------------------------------------------------------------
	// Expected output:
	//   After append: [Write tests Deploy app Review PR Fix bug]
	//   0: Write tests
	//   1: Deploy app
	//   2: Review PR
	//   3: Fix bug

	fmt.Println("--- Part 3: append & range ---")

	tasks := []string{"Write tests", "Deploy app", "Review PR"}

	// TODO: Append "Fix bug" to tasks (remember: reassign the result!)
	// tasks = append(tasks, "Fix bug")
	// fmt.Println("After append:", tasks)

	// TODO: Use a for-range loop to print each task with its index
	// for i, task := range tasks {
	//     fmt.Printf("%d: %s\n", i, task)
	// }

	_ = tasks // remove this line once you use tasks above

	fmt.Println()

	// -----------------------------------------------------------------
	// PART 4: Slice of structs (preview of Day 2)
	// -----------------------------------------------------------------
	// Expected output:
	//   {Buy milk false}
	//   {Learn Go false}
	//   {Walk dog false}
	//   Completed tasks: 0

	fmt.Println("--- Part 4: Slice of structs ---")

	type Task struct {
		Title string
		Done  bool
	}

	// TODO: Create a slice of Task structs with 3 tasks (all Done: false)
	// todos := []Task{
	//     {Title: "Buy milk"},
	//     {Title: "Learn Go"},
	//     {Title: "Walk dog"},
	// }

	// TODO: Print each task
	// for _, t := range todos {
	//     fmt.Println(t)
	// }

	// TODO: Count how many tasks have Done == true and print the count
	// count := 0
	// for _, t := range todos {
	//     if t.Done {
	//         count++
	//     }
	// }
	// fmt.Println("Completed tasks:", count)

	fmt.Println()

	// -----------------------------------------------------------------
	// PART 5: Challenge — write a function
	// -----------------------------------------------------------------
	// Write a function called `filterDone` that takes a []Task and returns
	// a new []Task containing only the tasks where Done == true.
	//
	// Then mark one task as done and verify your function works.
	//
	// Hint: start with an empty slice and append to it inside a loop.

	fmt.Println("--- Part 5: Challenge ---")
	// Your code here...

	fmt.Println()
	fmt.Println("=== Done! Uncomment the answer section below to check your work ===")
}

// =============================================================================
// ANSWERS — uncomment to check your work
// =============================================================================
//
// func filterDone(tasks []Task) []Task {
//     var result []Task
//     for _, t := range tasks {
//         if t.Done {
//             result = append(result, t)
//         }
//     }
//     return result
// }
//
// Full Part 1:
//   var scores [5]int
//   scores[2] = 99
//   fmt.Println("Array:", scores)
//   fmt.Println("After setting index 2:", scores)
//   fmt.Println("Length:", len(scores))
//
// Full Part 2:
//   tasks := []string{"Write tests", "Deploy app", "Review PR"}
//   fmt.Println("Tasks:", tasks)
//   fmt.Printf("Length: %d, Capacity: %d\n", len(tasks), cap(tasks))
//   fmt.Println("First two:", tasks[:2])
//
// Full Part 3:
//   tasks = append(tasks, "Fix bug")
//   fmt.Println("After append:", tasks)
//   for i, task := range tasks {
//       fmt.Printf("%d: %s\n", i, task)
//   }
//
// Full Part 4:
//   todos := []Task{
//       {Title: "Buy milk"},
//       {Title: "Learn Go"},
//       {Title: "Walk dog"},
//   }
//   for _, t := range todos {
//       fmt.Println(t)
//   }
//   count := 0
//   for _, t := range todos {
//       if t.Done { count++ }
//   }
//   fmt.Println("Completed tasks:", count)
