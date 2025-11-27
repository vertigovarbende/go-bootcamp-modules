package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("arithmetic operations")

	/*
		var a float32 = 10
		var b float32 = 3
		var result float32
	*/

	a := 10
	b := 3

	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Substraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	const pi float64 = 22 / 7.0
	fmt.Println(pi)

	// Overflow with signed integers
	var maxInt int64 = 9223372036854775807 // max value that int64 can hold
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	// Overflow with unsigned integers
	var uMaxInt uint64 = 9223372036854775807*2 + 1 // max value for uint64 type
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	// Underflow with floating point numbers
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)

}
