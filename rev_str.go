package main

import "fmt"

func main() {
	var str string
	fmt.Print("Enter a string: ")
	fmt.Scan(&str)

	// Reverse the string
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}

	fmt.Printf("The reversed string is: %s\n", reversedStr)
}
