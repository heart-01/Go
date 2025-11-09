package main

import "fmt"

func Defer() {

	defer fmt.Println("world")
	fmt.Println("hello")

}
