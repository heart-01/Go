package main

import "fmt"

func main() {
	tmp1 := make([]int, 3, 5) // len=3, cap=5
	tmp1[0] = 1
	showSlice(tmp1)

	var tmp2 []int // len=0, cap=0
	tmp2 = append(tmp2, 1)
	showSlice(tmp2)

	tmp3 := []int{1, 2, 3, 4, 5, 6} // len=6, cap=6
	showSlice(tmp3)                 // [1 2 3 4 5 6]
	// leading remove element in array
	tmp3 = tmp3[2:len(tmp3)]
	showSlice(tmp3) // [3 4 5 6]
	tmp3 = tmp3[1:2]
	showSlice(tmp3) // [4]
	// trailing remove element in array
	tmp3 = append(tmp3, 5, 6, 7, 8, 9, 10)
	tmp3 = tmp3[0 : len(tmp3)-1]
	showSlice(tmp3) // [4 5 6 7 8 9]

	tmp4 := []int{1, 2, 3, 4, 5, 6} // len=6, cap=6
	removeIndex(tmp4, 2)
	showSlice(tmp4) // [1 2 4 5 6]
}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func removeIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}
