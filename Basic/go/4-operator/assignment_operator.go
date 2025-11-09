package main

import "fmt"

func exampleAssignmentOperator() {
	var x int

	// Set value
	x = 5
	fmt.Println("x = 5 ->", x) // result: 5

	// Add value and set value
	x += 3
	fmt.Println("x += 3 ->", x) // result: 8

	// Subtract value and set value
	x -= 2
	fmt.Println("x -= 2 ->", x) // result: 6

	// Multiply value and set value
	x *= 4
	fmt.Println("x *= 4 ->", x) // result: 24

	// Divide value and set value
	x /= 2
	fmt.Println("x /= 2 ->", x) // result: 12

	// Divide and set remainder
	x %= 5
	fmt.Println("x %= 5 ->", x) // result: 2

	// AND bitwise and set value
	x &= 3
	fmt.Println("x &= 3 ->", x) // result: 2

	// OR bitwise and set value
	x |= 1
	fmt.Println("x |= 1 ->", x) // result: 3

	// XOR bitwise and set value
	x ^= 1
	fmt.Println("x ^= 1 ->", x) // result: 2

	// Shift left and set value
	x <<= 1
	fmt.Println("x <<= 1 ->", x) // result: 4

	// Shift right and set value
	x >>= 1
	fmt.Println("x >>= 1 ->", x) // result: 2
}
