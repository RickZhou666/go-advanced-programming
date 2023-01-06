package chap0

import (
	"fmt"
	"testing"
)

// https://www.programiz.com/golang/arrays
// An array is a collection of similar types of data.

// execute all tests
// $ go test -v --cover -count=1 ./...

// execute s04_arrays_test.go
// $ go test -v --cover -count=1 ./chap0/s04_arrays_test.go

// 1. Creating an array in Go

func TestCreatingArray(t *testing.T) {

	// 1.1 declare array variable of type integer
	// var array_variable = [size]datatype{elements of array}
	var arrayOfInteger = [5]int{1, 5}
	fmt.Println(arrayOfInteger)

	// 1.2 Declare an array of undefined size
	// var array_variable = [...]datatype{elements of array}
	// Here, if [] is left empty, it becomes a slice.
	// So [...] is a must if we want an undefined size array.
	var arrayOfString = [...]string{"Hello", "Rick"}
	fmt.Println(arrayOfString)

	// 1.3 Shorthand notation to declare an array
	arrayOfString2 := [...]string{"Hello", "Eve"}
	fmt.Println(arrayOfString2)
}

// 2. Accessing array elements in Golang
// In Go, each element in an array is associated with a number. The number is known as an array index.
// index start from 0
func TestAccessingArray(t *testing.T) {
	languages := [3]string{"Go", "Java", "C++"}

	fmt.Println("Access element at index 0: ", languages[0])
	fmt.Println("Access element at index 2: ", languages[2])
}

// 3. Initialize an Array in Golang
// We can also use index numbers to initialize an array.

func TestInitializeArray(t *testing.T) {
	// 3.1
	var arrayOfIntegers [3]int

	arrayOfIntegers[0] = 5
	arrayOfIntegers[1] = 10
	arrayOfIntegers[2] = 15
	fmt.Println(arrayOfIntegers)

	// 3.2 Initialize specific elements of an array
	// 		initialized the element at index 0 and 3
	// 		other index will be value of 0
	arrayOfIntegers2 := [5]int{0: 7, 3: 9}

	fmt.Println(arrayOfIntegers2)
}

// 4. Changing the array element in Go
func TestChangeArrayElement(t *testing.T) {
	weather := [3]string{"Rainy", "Sunny", "Cloudy"}

	weather[2] = "Stormy"
	fmt.Println(weather)
}

// 5. Find the length of an Array in Go
func TestFindLengthOfArray(t *testing.T) {
	var arrayOfIntegers = [...]int{1, 5, 8, 0, 3, 10}

	// find the length of array using len()
	length := len(arrayOfIntegers)

	fmt.Println("The length of array is: ", length)
}

// 6. Looping through an array in Go
func TestLoopArray(t *testing.T) {
	age := [...]int{12, 4, 5}
	for i := 0; i < len(age); i++ {
		fmt.Println(age[i])

	}
}

// 7. Multidimensional array in Golang
func TestMultidimensionArray(t *testing.T) {
	arrayInteger := [2][2]int{{1, 2}, {3, 4}}

	for i := 0; i < len(arrayInteger); i++ {
		for j := 0; j < len(arrayInteger[0]); j++ {
			fmt.Println(arrayInteger[i][j])
		}

	}
}
