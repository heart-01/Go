package main

import "fmt"

func main() {
	var myArr1 []int = []int{1, 2, 3, 4, 5}
	GetArray(myArr1)

	var myArr2 = []int{6, 7, 8, 9, 10}
	GetArray(myArr2)

	var myArr3 [3]int
	myArr3[0], myArr3[1], myArr3[2] = 11, 12, 13
	for index, item := range myArr3 {
		fmt.Println(index, item)
	}
}

func GetArray(arr []int) {
	for index, item := range arr {
		fmt.Println(index, item)
	}
}
