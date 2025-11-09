package main

import "fmt"

func main() {

	// Declare a struct
	type Person struct {
		name string
		age  int
	}

	// assign value to struct
	person1 := Person{"John", 30}
	fmt.Println(person1) // {John 30}

	// define an instance of struct
	var person2 Person

	// assign value to struct
	person2 = Person{
		name: "Jane",
		age:  25,
	}

	fmt.Println(person2) // {Jane 25}

}
