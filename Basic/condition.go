package main

import "fmt"

func main() {
	someValue := 9
	fnIfElse(someValue)
	fnSwitchCase()
}

func fnIfElse(someValue int) {
	if someValue == 10 {
		fmt.Println("a is 10")
	} else {
		fmt.Println("a is not 10")
	}

	if someValue > 10 || someValue < 20 {
		fmt.Println("a is greater than 10 or less than 20")
	} else {
		fmt.Println("a is not greater than 10 or less than 20")
	}

	if result := doSomething(); result == "OK" {
		fmt.Println("OK")
	} else {
		fmt.Println("Not OK")
	}
}

func doSomething() string {
	return "OK"
}

func fnSwitchCase() {
	index := 1
	switch index {
	case 0:
		fmt.Println("Zero")
		break
	case 1:
		fmt.Println("One")
		break
	case 2:
		fmt.Println("Two")
		break
	default:
		fmt.Println("Unknown Number")
		break
	}
}
