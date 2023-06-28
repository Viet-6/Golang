package functions

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

func areaOfSquare(a int) (edge int, area int) {
	edge = a

	area = edge * edge

	return
}

func FunctionExecute() {
	printHelloWorld()
	sayHi("V")
	fmt.Println(sum2(1, 1))

	a, b := 1, 2
	fmt.Println("before swap:")
	fmt.Println(a, b)
	a, b = swap(a, b)
	fmt.Println("after swap:")
	fmt.Println(a, b)

	edge, area := areaOfSquare(4)

	fmt.Printf("area of square with edge = %v is : %v", edge, area)
}
