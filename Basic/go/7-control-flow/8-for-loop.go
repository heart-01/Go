package main

import "fmt"

func Forloop() {

	for i := 1; i <= 10; i++ {
		if i == 5 {
			// break
			continue
		}
		fmt.Println(i)
	}

}
