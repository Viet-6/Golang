package main

import "fmt"

func printHelloWorld() {
	fmt.Println("Hello World")
}

func sayHi(name string) {
	fmt.Printf("Hi %v\n", name)
}

func sum2(a int, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func functionExecute() {
	printHelloWorld()
	sayHi("V")
	fmt.Println(sum2(1, 1))

	a, b := 1, 2
	fmt.Println("before swap:")
	fmt.Println(a, b)
	a, b = swap(a, b)
	fmt.Println("after swap:")
	fmt.Println(a, b)

}