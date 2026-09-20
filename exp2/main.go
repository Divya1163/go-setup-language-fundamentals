package main

import (
	"fmt"

	"go-lab/exp2/myutils"
)

func main() {

	fmt.Println("CUSTOM PACKAGE DEMONSTRATION")
	fmt.Println("-----------------------------")

	// String functions
	text := "Hello World"

	fmt.Println("Original string:", text)
	fmt.Println("Reversed string:", myutils.Reverse(text))
	fmt.Println("Number of vowels:", myutils.CountVowels(text))

	// Mathematical functions
	number := 5
	base := 2
	exponent := 5

	fmt.Println("Factorial of", number, ":", myutils.Factorial(number))
	fmt.Println(base, "raised to the power of", exponent, ":", myutils.Power(base, exponent))
}