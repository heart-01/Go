package main

import "fmt"

func main() {
	fn1()
	fn2("Hello, Go2")

	const str1, str2 = "Hello, ", "Go3"
	fmt.Println(sumString(str1, str2))

	const num1, num2 = 10, 20
	fmt.Printf("%d + %d = %d \n", num1, num2, sum(num1, num2))

	var num3, num4 int = swap(num1, num2)
	fmt.Printf("%d, %d => %d, %d \n", num1, num2, num3, num4)

	num5, num6, title := swapV2(num1, num2)
	fmt.Printf("%s %d, %d => %d, %d", title, num1, num2, num5, num6)
}

func fn1() {
	fmt.Println("Hello, Go1")
}

func fn2(msg string) {
	fmt.Println(msg)
}

func sumString(num1 string, num2 string) string {
	return num1 + num2
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func swap(num1 int, num2 int) (int, int) {
	return num2, num1
}

func swapV2(a, b int) (x, y int, title string) {
	x, y = b, a
	title = "SwapV2"
	return
}
