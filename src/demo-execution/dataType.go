package main

import (
	"fmt"
	"math"
	"math/bits"
)

/*
					Golang basic data types:
	1. bool

	2. string

	3. int int8 int16 int32 int64

	4. uint uint8 uint16 uint32 uint64 uintptr

	5. byte // alias for uint8

	6. rune // alias for int32 (represents Unicode)

	7. float32 float64

	8. complex64 complex128
							Advande

	- Zero values
	- Type conversion
	- constants
*/

func dataType() {

	var boolean bool = true // false
	fmt.Printf("boolean value: %v.\nboolean type: %T\n", boolean, boolean)
	fmt.Println("-----------------------------------------")

	var str string = "This is a string"
	fmt.Printf("str value: %v.\nstr type: %T\n", str, str)
	fmt.Println("-----------------------------------------")

	// int range
	// -9223372036854775808 -> 9223372036854775807
	fmt.Printf("int min: %v\n", math.MinInt)
	fmt.Printf("int max: %v\n", math.MaxInt)
	// -128 -> 127
	fmt.Printf("int8 min: %v\n", math.MinInt8)
	fmt.Printf("int8 max: %v\n", math.MaxInt8)
	// -32768 -> 32767
	fmt.Printf("int16 min: %v\n", math.MinInt16)
	fmt.Printf("int16 max: %v\n", math.MaxInt16)
	// -2147483648 -> 2147483647
	fmt.Printf("int32 min: %v\n", math.MinInt32)
	fmt.Printf("int32 max: %v\n", math.MaxInt32)
	// int64 range = int = -9223372036854775808 -> 9223372036854775807
	fmt.Printf("int64 min: %v\n", math.MinInt64)
	fmt.Printf("int64 max: %v\n", math.MaxInt64)
	fmt.Println("-----------------------------------------")

	// int bit
	// 64
	fmt.Printf("int bit: %v\n", bits.OnesCount(math.MaxUint))
	// 8
	fmt.Printf("int8 bit: %v\n", bits.OnesCount8(math.MaxUint8))
	// 16
	fmt.Printf("int16 bit: %v\n", bits.OnesCount16(math.MaxUint16))
	// 32
	fmt.Printf("int32 bit: %v\n", bits.OnesCount32(math.MaxUint32))
	// int64 = int = 64
	fmt.Printf("int64 bit: %v\n", bits.OnesCount64(math.MaxUint64))
	fmt.Println("-----------------------------------------")
}
