package main

import (
	"fmt"
)

func main() {

	var a int
	var b int
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	fmt.Printf("your choice \n1 addition \n2 subtraction \n3 multiplication \n4 division \n5 modulus \n6 exit\n")
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Printf("Addition of %d and %d is %d\n", a, b, a+b)
	case 2:
		fmt.Printf("Subtraction of %d and %d is %d\n", a, b, a-b)
	case 3:
		fmt.Printf("Multiplication of %d and %d is %d\n", a, b, a*b)
	case 4:
		if b != 0 {
			fmt.Printf("Division of %d and %d is %d\n", a, b, a/b)
		} else {
			fmt.Println("Division by zero is not allowed")
		}
	case 5:
		fmt.Printf("Modulus of %d and %d is %d\n", a, b, a%b)
	case 6:
		fmt.Println("Exiting the program")
	default:
		fmt.Println("Invalid choice")
	}

}
