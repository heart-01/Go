package main

import "fmt"

func main() {
	// Explicit Declaration
	var x string = "Hello, Go1"
	x = "Hello, Go2"
	const y string = "Hello, Go3"

	// Implicit Declaration
	z := "Hello, Go4"
	z = "Hello, Go5"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
