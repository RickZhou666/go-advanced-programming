package ch3

import (
	"fmt"
	"testing"
)

// https://www.programiz.com/golang/slice
// 1. Create slice in Golang
func TestCreatSlice(t *testing.T) {
	// numbers 			- the name of the slice
	// int 				- slice only includes integer numbers
	// {1, 2, 3, 4, 5} 	- elements of the slice
	numbers := []int{1, 2, 3, 4, 5}
	// we have left the [] notation empty.
	// This is because we can add elements to a slice that will change its size.

	// If we provide size inside the [] notation, it becomes an array.

	fmt.Println("Numbers: ", numbers)
}
