// Convert from integer to boolean (displayed as integer because Go does not support direct conversion from int to bool)

package main

import "fmt"

func TypeConversion5() {
	var i int = 1
	var b bool = (i != 0)
	fmt.Printf("i: %d, b: %t\n", i, b)
	// result: i: 1, b: true
}
