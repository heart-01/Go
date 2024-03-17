package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		defer printFinish(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Start: ", i)
	}
}

func printFinish(i int) {
	fmt.Println("Finish: ", i)
}
