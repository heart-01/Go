package main

import (
	"fmt"
	uuid "github.com/google/uuid"
	sample "github.com/heart-01/Go/func"
)

func main() {
	sample.Hello()

	id := uuid.New()
	fmt.Println("Generated UUID:", id)
}
