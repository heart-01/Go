// Convert from integer to decimal
package main

import (
	"fmt"
)

func TypeConversion1() {
	var i int = 42
	var f float64 = float64(i)
	fmt.Printf("i: %d, f: %f\n", i, f)
	// result: i: 42, f: 42.000000
}
