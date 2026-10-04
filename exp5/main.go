package main

import "fmt"

func modifyVal(n *int) {
	*n = 100
}

// Struct
type Student struct {
	Name string
	Age  int
}

func main() {

	//Referencing and dereferencing
	value := 50

	fmt.Println("Original value:", value)
	fmt.Println("Address of value:", &value)

	ptr := &value
	fmt.Println("Value using pointer (*ptr):", *ptr)

	//Passing pointer to a function
	fmt.Println("\nBefore function call:", value)

	modifyVal(&value)

	fmt.Println("After function call:", value)

	//Allocate a struct using new()
	studentPtr := new(Student)

	studentPtr.Name = "Divya"
	studentPtr.Age = 23

	fmt.Println("\nStudent details:")
	fmt.Println("Name:", studentPtr.Name)
	fmt.Println("Age:", studentPtr.Age)

	// Modify field through pointer
	studentPtr.Age = 24

	fmt.Println("Updated age:", studentPtr.Age)
}