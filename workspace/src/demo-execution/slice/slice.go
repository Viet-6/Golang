package slice

import "fmt"

func Slice() {
	// Slice type is an abstraction built on top of array type
	var first_slice []int
	fmt.Println(first_slice) // []
	var second_slice = []int{1, 2, 3, 4}
	fmt.Println(second_slice)
	fmt.Println(len(second_slice))

	// declare slice from an array
	arr := [4]int{100, 200, 300, 400}
	slice1 := arr[1:3]

	fmt.Println(slice1)

	slice2 := arr[:]
	fmt.Println(slice2)

	slice3 := arr[3:]
	fmt.Println(slice3)

	slice4 := arr[:2]
	fmt.Println(slice4)

	// delcare slice from a slice
	new_slice := []int{99, 88, 77, 66}
	nslice1 := new_slice
	fmt.Println(nslice1)

	// length & capacity
	// 1. Length is the size of the slice
	// 2. capacity is the size of the underlying array start from the index defined in slice
	fmt.Println(len(slice1)) // 2
	fmt.Println(cap(slice1)) // 3

	// method working with slice: make, copy, append

	// make: create a slice with specify length and capacity
	// Note: if not define capacity it capacity will equal to length
	slice6 := make([]int, 2, 6)
	fmt.Println(slice6)
	fmt.Println(len(slice6))
	fmt.Println(cap(slice6))

	// append: add element into slice
	// capacity will be x2 when the number of element in slice is over the defined capacity
	slice6 = append(slice6, 123)
	slice6 = append(slice6, 123)
	slice6 = append(slice6, 123)
	slice6 = append(slice6, 123)
	slice6 = append(slice6, 123)
	fmt.Println(slice6)
	fmt.Println(len(slice6))
	fmt.Println(cap(slice6))

	slice6[0], slice6[1] = 1, 2

	// copy:
	slice7 := make([]int, 2)
	new_s := copy(slice7, slice6) // copy() will return the number of copied element
	fmt.Println("____________________")
	fmt.Println(new_s)
	fmt.Println(slice7)

	// slice trick: remove an element in slice
	// The way to achive this to divide slice in to 2 small slice and concat 2 slice together
	// Ex: remove "C"
	strs := []string{"A", "B", "C", "D"}
	strs = append(strs[:2], strs[3:]...)
	fmt.Println(strs)

	// slice -> reference type
	nArr := [4]int{1, 2, 3, 4}
	slice10 := nArr[:]

	slice10[0] = 100
	fmt.Println(nArr)
	fmt.Println(slice10)

}
