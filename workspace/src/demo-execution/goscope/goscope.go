package goscope

import "fmt"

/*
	There are 3 level of scope in Go
	block -> package -> universe
*/

// package scope can be access anywhere as long as it in the same package
var global = 22

func Goscope() {
	// Block scope is the code inside {}

	{
		// Variable declare inside a block can only accessable inside the block
		// Variable declare inside the block is not relate to outside the block
		var x = 2
		var integer = 10

		fmt.Println(x, integer, global)
	}

	var integer = 1

	fmt.Println(integer)
	// fmt.Println(x)
	fmt.Println(global)
}

// universe scope is allow the variables or functions can be access outside of the package
// The way is achive this is make sure the first character is Capitalize
var Universe = "Universe"
