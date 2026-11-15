
package main

import (
	"fmt"
	"math"
	"strings"
)

// Function to perform division
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

// Function to calculate square root
func SquareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("cannot calculate square root of a negative number")
	}
	return math.Sqrt(n), nil
}

// Function to validate input
func ValidateInput(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("input cannot be empty")
	}
	return input, nil
}

func main() {
	fmt.Println("Error Handling in Go")

	// Test division
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division Result:", result)
	}

	// Test square root
	result, err = SquareRoot(25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square Root:", result)
	}

	// Test input validation
	value, err := ValidateInput("Hello Go")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Valid Input:", value)
	}

	// Demonstrate error cases
	_, err = Divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = SquareRoot(-9)
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = ValidateInput("")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
