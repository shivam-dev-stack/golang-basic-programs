package main

import (
	"fmt"
)

func main() {

	var arr = []int{1, 6, 3, 11, 5, 25, 7, 8, 12, 10}
	var max int = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Printf("The maximum value in the array is: %d\n", max)
}
