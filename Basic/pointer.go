package main

import "fmt"

func main() {
	fnEx1()
	fnEx2()
}

func fnEx1() {
	i := 42
	p := &i

	fmt.Println("I = ", *p)
	fmt.Println("Address of I = ", p)

	*p = *p / 2

	fmt.Println("I = ", *p)
}

func fnEx2() {
	msg := "Hello, World!"
	var msgPointer *string = &msg
	fmt.Println("Message = ", msg)
	fmt.Println("Address of Message = ", msgPointer)
	fmt.Println("Value of Message = ", *msgPointer)

	changeMessage(&msg, "Hello Pointer1")
	fmt.Println(msg)

	changeMessage(msgPointer, "Hello Pointer2")
	fmt.Println(msg)

	changeMessage(msgPointer, "Hello Pointer3")
	fmt.Println(*msgPointer)
}

func changeMessage(pointer *string, newMessage string) {
	*pointer = newMessage
}
