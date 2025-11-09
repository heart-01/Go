package main

import "fmt"

func main() {
	// Create two variables
	i, j := 42, 2701

	// Create a pointer variable p and set it to the address of i
	p := &i // address of i like this: 0xc0000160e0 = value of p

	// Read value p through pointer
	fmt.Println(*p) // 42

	// Change value i through pointer
	*p = 21
	fmt.Println(i) // 21

	// Create a pointer variable p and set it to the address of j
	p = &j // address of j like this: 0xc0000160e8 = value of p

	// Change value j through pointer
	*p = *p / 37
	fmt.Println(j) // 73
}
