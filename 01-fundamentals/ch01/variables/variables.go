package main

import "fmt"

// If we want to declare global variable, we need to use 'var' keyword not ':='
var middleName string = "Cane"

func main() {
	// Default values
	// Numeric Types -> 0
	// Boolean Types -> false
	// String Type -> ""
	// Pointers, slices, maps, functions and structs -> nil

	fmt.Println("variables")

	var age int
	var name string = "John"
	// It is optional if we want to declare the variable as a String
	var name2 = "Doe"
	count := 10
	lastName := "Smith"

	fmt.Println(age, name, name2, count, lastName)
	fmt.Println(middleName)
	printName()

}

func printName() {
	firstName := "Micheal"
	fmt.Println(firstName)
}
