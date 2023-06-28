package functions

import "fmt"

/*
	Variadic function allow passing an infinite number of arguements with only define it name once
	It also allow to pass a slice into function
*/

func addItems(n int, list ...int) {
	// Variadic function convert the parameters passing after n to a slice and assign it to list
	// 1, 2, 3, 4, 5 -> []int{1, 2, 3, 4, 5}
	list = append(list, n)
	fmt.Println(list)
}

func VariadicExecute() {
	addItems(100, 1, 2, 3, 4, 5)

	// to pass an slice into Variadic function add ... after the slice
	slice := []int{-1, -2, -3}
	addItems(-6, slice...)
}
