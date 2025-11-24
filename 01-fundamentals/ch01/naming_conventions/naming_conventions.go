package main

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	fmt.Println("naming conventions")
	const MAXRETRIES = 5

	employeeId := 1001
	fmt.Println("EmployeeID: ", employeeId)
}
