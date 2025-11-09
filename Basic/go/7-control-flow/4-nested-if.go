package main

import "fmt"

func Nestestif() {

	number1 := 12 // 12, 22, 20
	number2 := 20

	if number1 >= number2 {
		if number1 == number2 {
			fmt.Printf("Result: %d == %d", number1, number2)

		} else {
			fmt.Printf("Result: %d > %d", number1, number2)
		}

	} else {
		fmt.Printf("Result: %d < %d", number1, number2)
	}
}
