package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	countInstance1 := intSeq()
	fmt.Println(countInstance1()) // 1
	fmt.Println(countInstance1()) // 2
	fmt.Println(countInstance1()) // 3
	fmt.Println(countInstance1()) // 4
	fmt.Println(countInstance1()) // 5

	countInstance2 := intSeq()
	fmt.Println(countInstance2()) // 1
	fmt.Println(countInstance2()) // 2
	fmt.Println(countInstance2()) // 3

	fmt.Println(intSeq()()) // 1
}
