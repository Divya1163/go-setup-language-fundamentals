package main

import "fmt"

const minValue = -1000
const maxValue = 1000

func getIntegerInput(message string) int {
	var num int

	for {
		fmt.Print(message)

		_, err := fmt.Scan(&num)

		if err != nil {
			fmt.Println("Invalid input! Please enter a numeric integer.")

			var discard string
			fmt.Scan(&discard)

			continue
		}

		if num < minValue || num > maxValue {
			fmt.Printf("Out of range! Enter a value between %d and %d.\n",
				minValue, maxValue)
			continue
		}

		return num
	}
}

func getFloatInput(message string) float64 {
	var num float64

	for {
		fmt.Print(message)

		_, err := fmt.Scan(&num)

		if err != nil {
			fmt.Println("Invalid input! Please enter a numeric value.")

			var discard string
			fmt.Scan(&discard)

			continue
		}

		if num < minValue || num > maxValue {
			fmt.Printf("Out of range! Enter a value between %d and %d.\n",
				minValue, maxValue)
			continue
		}

		return num
	}
}