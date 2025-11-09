// Convert from boolean to integer (represented as integer because Go does not support direct bool to int conversion).

package main

import "fmt"

func TypeConversion6() {
	var b bool = true
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	fmt.Printf("b: %t, i: %d\n", b, i)
	// result: b: true, i: 1
}
