package main

import "fmt"

func Ifelseifcondition() {

	number1 := 20
	number2 := 20

	if number1 == number2 {
		fmt.Printf("Result: %d == %d", number1, number2)
	} else if number1 > number2 {
		fmt.Printf("Result: %d > %d", number1, number2)
	} else {
		fmt.Printf("Result: %d < %d", number1, number2)
	}
}
