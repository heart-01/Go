package main

import "fmt"

func main() {
	fnFor(10)
	fnWhile(5)

	courses := []string{"Go", "Java", "Python", "C++"}
	fnForeach(courses)

	fnWhileUsingBreak()
}

func fnFor(count int) {
	for index := 0; index < count; index++ {
		fmt.Printf("For: %d \n", index)
	}
}

func fnWhile(count int) {
	index := 0
	for index < count {
		fmt.Printf("While: %d \n", index)
		index++
	}
}

func fnForeach(data []string) {
	for index, item := range data {
		fmt.Printf("Index: %d, Value: %s \n", index, item)
	}
}

func fnWhileUsingBreak() {
	index := 0
	for true {
		if index > 5 {
			break
		}
		fmt.Printf("WhileBreak-Index %d \n", index)
		index++
	}
}
