package main

import "fmt"

func exampleCompareOperator() {
	var x bool

	// equal to
	x = (5 == 3)
	fmt.Println("5 == 3 =", x) // result: false

	// not equal to
	x = (5 != 3)
	fmt.Println("5 != 3 =", x) // result: true

	// greater than
	x = (5 > 3)
	fmt.Println("5 > 3 =", x) // result: true

	// less than
	x = (5 < 3)
	fmt.Println("5 < 3 =", x) // result: false

	// greater than or equal to
	x = (5 >= 3)
	fmt.Println("5 >= 3 =", x) // result: true

	// less than or equal to
	x = (5 <= 3)
	fmt.Println("5 <= 3 =", x) // result: false
}
