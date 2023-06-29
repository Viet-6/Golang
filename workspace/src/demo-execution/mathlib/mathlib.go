package mathlib

import (
	"fmt"
	"math"
)

func Mathlib() {

	// Some common math function

	// math.Ceil() return the least integer that greater or equal to x
	fmt.Println(math.Ceil(3.001))
	fmt.Println(math.Ceil(3.999))

	// math.Min() return min of x and y
	fmt.Println(math.Min(100, 0))

	// math.Max() return max of x and y
	fmt.Println(math.Max(100, 0))

	// math.Abs() return absolute value
	fmt.Println(math.Abs(-1.9))

	// math.Sqrt return square root of x
	fmt.Println(math.Sqrt(100))

	// math.Pow return x**y
	fmt.Println(math.Pow(2, 4))
}
