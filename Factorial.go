package main

import "fmt"

func main() {

	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	var f int = 1

	if number == 0 || number == 1 {
		fmt.Printf("The factorial of %d is %d\n", number, 1)
	} else {
		for i := 1; i <= number; i++ {
			f = f * i
		}
		fmt.Printf("The factorial of %d is %d\n", number, f)
	}

	// Recursive function to calculate factorial
	fmt.Printf("The factorial of %d is %d\n", number, factorial(number))

}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
