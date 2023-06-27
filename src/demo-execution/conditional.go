package main

import "fmt"

func conditional() {
	ifCondition()
	switchCondition()
}

func switchCondition() {
	fmt.Println("Switch condition: ")
	n := 10
	fmt.Println("Basic switch:")

	switch n {
	case 1, 2, 3, 4, 5:
		fmt.Println("n <= 5")
	case 10, 11, 12, 13:
		fmt.Println("n >= 10")
	default:
		fmt.Printf("undefined")
	}

	fmt.Println("Switch with condition:")

	number := 3

	switch {
	case number < 5:
		fmt.Println("n <= 5")
	case number > 10:
		fmt.Println("n >= 10")
	default:
		fmt.Printf("undefined")
	}

	fmt.Println("switch fallThrough:")

	new_num := 11

	switch new_num {
	case 5:
		fmt.Println("number = 5")
		fallthrough
	case 11:
		fmt.Println("number = 11")
		fallthrough
	case 12:
		fmt.Println("number = 12")
		fallthrough
	default:
		fmt.Println("not know")
	}

	fmt.Println("switch with Goto:")

	another_num := 1

	switch another_num {
	case 1:
		fmt.Print("this num is equal to: ")
		goto nextNumber
		fmt.Print(1, "\n")
	nextNumber:
		fmt.Print(2, "\n")
	default:
		fmt.Print(nil, "\n")
	}

	fmt.Println("_________________________________")
}

func ifCondition() {
	fmt.Println("If condition: ")
	// if condition { // code }
	if 1 == (2 - 1) {
		fmt.Println("1 is equal to (2 - 1)")
	}

	// if condition { // code } else { // code }
	if 2 != (2 / 1) {
		fmt.Println("2 is not equal to (2 / 1)")
	} else {
		// In Go else statement can only write right after the close backet (})
		// It will be an error if the else statement is start in the new line
		fmt.Println("2 is equal to (2 / 1)")
	}

	// if statement; condition { // code }
	if a := 100; (a % 2) == 0 {
		fmt.Printf("%v is even number.\n", a)
	}

	fmt.Println("_________________________________")
}
