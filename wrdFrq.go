package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Print("Enter a string: ")
	fmt.Scan(&str)

	var char string
	fmt.Print("Enter a character: ")
	fmt.Scan(&char)

	var count int = 0

	// Count the frequency of each character in the string
	for i := 0; i < len(str); i++ {
		if string(str[i]) == char {
			count++
		}
	}

	fmt.Printf("The frequency of '%s' in the string is: %d\n", char, count)

}
