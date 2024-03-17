package employee

import "fmt"

type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

var employeeInstance *employee

// Singleton pattern
func Init(firstname string, lastname string, totalleaves int, leavestaken int) *employee {
	if employeeInstance == nil {
		employeeInstance = &employee{
			FirstName:   firstname,
			LastName:    lastname,
			TotalLeaves: totalleaves,
			LeavesTaken: leavestaken,
		}
	}
	return employeeInstance
}

func (employee employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", employee.FirstName, employee.LastName, employee.TotalLeaves-employee.LeavesTaken)
}
