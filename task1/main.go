package main

import "fmt"

func main() {
	var int1 int = 20
	var int2 int = 5

	var float1 float64 = 10.5
	var float2 float64 = 2.5

	// Integer operations
	fmt.Println("Integer Operations:")
	fmt.Println("Addition:", int1+int2)
	fmt.Println("Subtraction:", int1-int2)
	fmt.Println("Multiplication:", int1*int2)

	// Floating-point operations
	fmt.Println("\nFloating-Point Operations:")
	fmt.Println("Addition:", float1+float2)
	fmt.Println("Subtraction:", float1-float2)
	fmt.Println("Multiplication:", float1*float2)
}

