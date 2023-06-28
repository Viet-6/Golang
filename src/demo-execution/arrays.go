package main

import "fmt"

func arrays() {
	var firstArr [4]int
	fmt.Println("Default array value: ", firstArr)

	var secondArr = [4]int{1, 2, 3, 4}
	fmt.Println("secondArr value: ", secondArr)

	// var thirdArrs [2]string = {"V", "D"}
	// Declare array with undefined range(length).
	// The range of the array will be fix with the value in it when declare
	var thirdArr = [...]int{100, 200, 300}
	fmt.Println("thirdArr value: ", thirdArr)

	// In Go you have 2 ways to loop through an array
	// 1, using len() function
	for i := 0; i < len(thirdArr); i++ {
		fmt.Println(thirdArr[i])
	}

	// 2. using range keyword. (similar like foreach)
	for index, value := range secondArr {
		fmt.Printf("The value of index %v is %v\n", index, value)
	}

	// With the second way you only want to works with one of the two values index or value you can replace it with _
	// This was called blank identifier
	// Note: `_, _` is not allowed
	for _, value := range thirdArr {
		fmt.Printf("value = %v\n", value)
	}

	// 2-dimensional arrays
	// col row
	matrix := [3][2]int{
		{1, 2},
		{2, 3},
		{3, 4},
	}

	fmt.Println(matrix)
	fmt.Println("column length = ", len(matrix)) // equal col = 3
	fmt.Println("row length = ", len(matrix[0])) // equal row = 2

	// loop through 2-dimensional array
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

}
