package main

import "fmt"

func main() {

	// PART 1: SLICE OPERATIONS

	fmt.Println("SLICE OPERATIONS")

	// Initial slice
	students := []string{"Alice", "Bob", "Charlie"}
	fmt.Println("Initial slice:", students)

	// Add an element
	students = append(students, "David")
	fmt.Println("After adding David:", students)

	// Remove element by index
	// Removing element at index 1 (Bob)
	index := 1
	students = append(students[:index], students[index+1:]...)
	fmt.Println("After removing element at index 1:", students)

	// Update an element
	// Changing Charlie
	students[1] = "James"
	fmt.Println("After updating index 1:", students)

	
	// PART 2: MAP OPERATIONS

	fmt.Println()
	fmt.Println("MAP OPERATIONS")

	// Initial map
	marks := map[string]int{
		"Math":    85,
		"Science": 90,
		"English": 78,
	}

	fmt.Println("Initial map:", marks)

	// Insert a key-value pair
	marks["Computer"] = 95
	fmt.Println("After inserting Computer:", marks)

	// Delete a key-value pair
	delete(marks, "English")
	fmt.Println("After deleting English:", marks)

	// Lookup a value
	subject := "Math"
	mark, exists := marks[subject]

	if exists {
		fmt.Println("Lookup -", subject, "marks:", mark)
	} else {
		fmt.Println(subject, "not found")
	}
}