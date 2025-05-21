package main

import (
	"fmt"
)

func main() {

	var number int = 50

	// Check if the number is even or odd
	if number%2 == 0 {
		fmt.Printf("The number %d is even.\n", number)
	} else {
		fmt.Printf("The number %d is odd.\n", number)
	}
}
