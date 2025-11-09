package main

import "fmt"

func exampleMathOperator() {
	var x int

	// Addition
	x = 5 + 3
	fmt.Println("5 + 3 =", x) // result: 8

	// Subtraction
	x = 5 - 3
	fmt.Println("5 - 3 =", x) // result: 2

	// Multiplication
	x = 5 * 3
	fmt.Println("5 * 3 =", x) // result: 15

	// Division
	x = 10 / 2
	fmt.Println("10 / 2 =", x) // result: 5

	// Modulo
	x = 10 % 3
	fmt.Println("10 % 3 =", x) // result: 1

	// Increment
	x = 5
	x++
	fmt.Println("5++ =", x) // result: 6

	// Decrement
	x = 5
	x--
	fmt.Println("5-- =", x) // result: 4
}
