package main

import (
	"struct_workshop/employee"
)

func main() {
	e := employee.Init("Admin", "Test", 30, 20)
	e.LeavesRemaining()
}
