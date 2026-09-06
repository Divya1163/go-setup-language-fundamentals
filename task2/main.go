package main

import "fmt"

func main() {

	fmt.Println("--Simple Calculator--")

	// Integer calculations
	fmt.Println("\n--- Integer Calculator ---")

	int1 := getIntegerInput("Enter first integer: ")
	int2 := getIntegerInput("Enter second integer: ")

	fmt.Println("\nInteger Results:")
	fmt.Println("Addition:", int1+int2)
	fmt.Println("Subtraction:", int1-int2)
	fmt.Println("Multiplication:", int1*int2)

	// Floating-point calculations
	fmt.Println("\n--- Floating-Point Calculator ---")

	float1 := getFloatInput("Enter first floating-point number: ")
	float2 := getFloatInput("Enter second floating-point number: ")

	fmt.Println("\nFloating-Point Results:")
	fmt.Println("Addition:", float1+float2)
	fmt.Println("Subtraction:", float1-float2)
	fmt.Println("Multiplication:", float1*float2)

	fmt.Println("\n-- Calculator Completed --")
}