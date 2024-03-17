package main

import "fmt"

func main() {
	numbers := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("", numbers["two"])

	var colors = make(map[string]string)
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"
	colors["blue"] = "#0000ff"
	fmt.Println("", colors["green"])

	courses := make(map[string]map[string]int)
	courses["Android"] = make(map[string]int)
	courses["Android"]["price"] = 200
	courses["iOS"] = make(map[string]int)
	courses["iOS"]["price"] = 300
	fmt.Println("", courses)
}
