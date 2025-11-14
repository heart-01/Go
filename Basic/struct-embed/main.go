package main

import (
	"fmt"
	"struct-embed/test1"
	"struct-embed/test2"
)

func main() {

	// Create struct from test1
	p1 := test1.Profile{
		Firstname: "Hi",
		Lastname:  "Dev",
	}

	// Use struct from test2 that embeds test1.Profile
	p2 := test2.Profile{
		Profile: &p1,
	}

	// Method from test1
	fmt.Println("Firstname:", p2.GetFirstname())

	// Method from test2
	fmt.Println("Lastname:", p2.GetLastname())

	// ❌ Not allowed: p1 does not have GetLastname()
	// fmt.Println(p1.GetLastname()) // error
}
