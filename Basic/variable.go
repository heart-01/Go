package main

import "fmt"

var count int = 10

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

	run()
}

func run() {
	var tem int = 20
	count++
	fmt.Println(count)
	fmt.Println(tem)
}
