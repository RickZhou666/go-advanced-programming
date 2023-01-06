package chap0

import (
	"fmt"
	"testing"
)

// https://www.programiz.com/golang/struct
// A struct is used to store variables of different data types

// 1. Declare Go Struct
type StructureName struct {
	// structure definition
}

// 2. Struct instance
// var person1 Person

func TestDeclareAStruct(t *testing.T) {
	// decalre a struct
	type Person struct {
		name string
		age  int
	}

	// assign value to struct while creating an instance
	person1 := Person{"John", 25}
	fmt.Println(person1)

	// define an instance
	var person2 Person

	// assign value to struct variables
	person2 = Person{
		name: "Sara",
		age:  29,
	}
	fmt.Println(person2)
}

// 3. Access a struct in Golang
// We can access individual elements of a struct using the struct instances
// we have used the . (dot) symbol to access the property of a struct.
func TestAccessStruct(t *testing.T) {
	// declare a struct
	type Rectangle struct {
		length  int
		breadth int
	}

	// declare instance rect1 and defining the struct
	rect := Rectangle{22, 12}

	// access the length of the struct
	fmt.Println("Length: ", rect.length)

	// access the breadth of the struct
	fmt.Println("Breadth: ", rect.breadth)

	area := rect.length * rect.breadth
	fmt.Println("Area: ", area)
}

// 4. Function inside a Struct in Go
// Go also allows us to create a function inside a struct. It treats function as a field of struct.

func TestFunctionInsideStruct(t *testing.T) {
	// initialize the function Rectangle
	type Rectangle func(int, int) int

	// create structure
	type reactanglePara struct {
		length  int
		breadth int
		color   string

		// function as a field of struct
		rect Rectangle
	}

	result := reactanglePara{
		length:  10,
		breadth: 20,
		color:   "Red",
		rect: func(length int, breadth int) int {
			return length * breadth
		},
	}

	fmt.Println("Color of Rectangle: ", result.color)
	fmt.Println("Area of Rectangle: ", result.rect(result.length, result.breadth))
}
