package main

import "fmt"

// Define Person struct
type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

// Method to read data into Person
func (p *Person) readData() {
	fmt.Print("Enter Name: ")
	fmt.Scan(&p.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&p.Age)

	fmt.Print("Enter Job: ")
	fmt.Scan(&p.Job)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&p.Salary)
}

// Method to display Person details
func (p Person) printData() {
	fmt.Println("\n--- Person Details ---")
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Job:", p.Job)
	fmt.Printf("Salary: %.2f\n", p.Salary)
}

func main() {

	var person1 Person
	fmt.Println("Enter details for Person 1:")
	person1.readData()

	var person2 Person
	fmt.Println("\nEnter details for Person 2:")
	person2.readData()

	person1.printData()
	person2.printData()
}