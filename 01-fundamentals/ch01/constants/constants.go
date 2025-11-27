package main

import "fmt"

// un-typed constants
const pi = 3.14
const GRAVITY = 9.81

func main() {
	fmt.Println("constants")

	// typed-constant
	const days int = 7

	// multiple constants - const block
	const (
		monday        = 1
		tuesday       = 2
		wednesday     = 3
		thursday  int = 4
		friday        = 5
		saturday      = 6
		sunday        = 7
	)

	fmt.Println(pi, GRAVITY)
	fmt.Println(days)
	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

}
