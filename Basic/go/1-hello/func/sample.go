package sample

import (
	"fmt"
	utils "github.com/heart-01/Go/func/internal"
)

func Hello() {
	fmt.Println("Hello, from Sample")
	result := utils.Add(10, 20)
	fmt.Println("Addition Result:", result)
}
