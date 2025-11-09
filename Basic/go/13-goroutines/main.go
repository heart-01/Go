package main

import (
	"fmt"
	"time"
)

// create function goroutine
func printNumbers() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Microsecond) // Sleep for 100 microseconds
	}
}

func main() {

	// Call printNumbers function in goroutine
	go printNumbers()

	// Main function in Goroutine will run concurrently with main function
	for i := 'A'; i <= 'E'; i++ {
		fmt.Println(string(i))
		time.Sleep(150 * time.Microsecond) // Sleep for 150 microseconds
	}

}
