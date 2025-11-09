// Convert from integer to string

package main

import (
	"fmt"
	"strconv"
)

func TypeConversion3() {
	var i int = 42
	var s string = strconv.Itoa(i)
	fmt.Printf("i: %d, s: %s\n", i, s)
	// result: i: 42, s: 42
}
