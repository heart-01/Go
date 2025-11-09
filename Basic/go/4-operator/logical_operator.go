package main

import "fmt"

func exampleLogicalOperator() {
	var x bool

	// AND
	x = (true && false)
	fmt.Println("true && false =", x) // result: false

	// OR
	x = (true || false)
	fmt.Println("true || false =", x) // result: true

	// NOT
	x = !true
	fmt.Println("!true =", x) // result: false
}
