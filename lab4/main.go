
package main

import (
	"fmt"
	"time"
)

func calculateSquares(ch chan string) {
	fmt.Println("Square Goroutine Started")
	for i := 1; i <= 5; i++ {
		result := i * i
		ch <- fmt.Sprintf("Square of %d = %d", i, result)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Square Goroutine Completed")
}

func calculateCubes(ch chan string) {
	fmt.Println("Cube Goroutine Started")
	for i := 1; i <= 5; i++ {
		result := i * i * i
		ch <- fmt.Sprintf("Cube of %d = %d", i, result)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Cube Goroutine Completed")
}

func calculateFibonacci(ch chan string) {
	fmt.Println("Fibonacci Goroutine Started")
	a, b := 0, 1

	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("Fibonacci number %d = %d", i, a)
		a, b = b, a+b
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Fibonacci Goroutine Completed")
}

func main() {
	ch := make(chan string)

	fmt.Println("Main function started")

	go calculateSquares(ch)
	go calculateCubes(ch)
	go calculateFibonacci(ch)

	// Collect 15 results from the channel
	for i := 1; i <= 15; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("All results collected. Main function completed.")
}
