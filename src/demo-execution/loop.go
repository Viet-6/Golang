package main

import "fmt"

func looping() {

	// for init; condition; post { // code }
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	// Golang dont have while statement but you can simulated it with for
	// for ; condition; { // code }
	// The semi-colom is the optional so you can remove it like this: for condition { // code }
	fmt.Println("--------------------------------")
	f := 5
	for f >= 0 {
		fmt.Print(f, " ")
		f--
	}
	fmt.Println()
	fmt.Println("--------------------------------")

	// infinite loop
	// for {
	// 	fmt.Println("looping ...")
	// }

	// Golang also support init multiple variables on for statement
	for i, j := 1, 2; i < 10 && j < 10; i, j = i+1, j+1 {
		fmt.Printf("i = %v; j = %v \t", i, j)
	}
	fmt.Println()
	fmt.Println("--------------------------------")
}
