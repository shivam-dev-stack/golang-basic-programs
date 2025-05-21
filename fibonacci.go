package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	a, b := 0, 1

	for i := 0; i < number; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()
}
