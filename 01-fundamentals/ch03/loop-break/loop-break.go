package main

import "fmt"

func main() {

	fmt.Println("loops")

	// Simple iteration over a range
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// iterate over collection
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		// d is specific to numbers
		fmt.Printf("Index: %d, Value: %d\n", index, value)
		// v is a general value
		// fmt.Printf("Index: %v, Value: %v\n", index, value)
	}

	// break and continue keywords
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd Number:", i)
		if i == 5 {
			break
		}
	}

}
