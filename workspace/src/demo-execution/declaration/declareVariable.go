package declaration

import "fmt"

func DeclareVariable() {
	var number int
	number = 100

	fmt.Println(number)

	var new_number int = 100

	fmt.Println(new_number)

	var number1, number2 int
	fmt.Println("Default value of variable if no assign value:")
	fmt.Println(number1, number2)

	number1 = 100
	number2 = 200
	fmt.Println("value after assigned:")
	fmt.Println(number1, number2)

	var number3, number4 int = 300, 400
	// var str string, n int = "str", 100

	fmt.Println(number3, number4)

	var (
		name   string
		age    int
		gender bool
	)

	name = "V"
	age = 23
	gender = true

	fmt.Println(name, age, gender)

	var (
		name1 string = "V2"
		age1  int    = 100
	)

	fmt.Println(name1, age1)

	var no_define_type = 1
	var number10, number23 = 100, "text"

	fmt.Println(no_define_type, number10, number23)

	new_name := "text"
	number100 := 100
	strs, boolean := "text", true

	fmt.Println(new_name, number100, strs, boolean)

	// Pointer
	x := 1
	var y int = 2
	pt := &x
	var px *int = &y

	fmt.Println(pt)
	fmt.Println(*pt)

	fmt.Println(px)
	fmt.Println(*px)

	// decalre a type
	type myInt int
	type myNewInt int
	var z myInt = 99
	var c myNewInt = 99
	fmt.Println(z)
	fmt.Println(c)

	// Even if the underlying type of myInt and myNewInt is int
	// you still can not assign or compare above 2 value without parse it to the same type (myNewInt or myInt)
	// fmt.Println(c == z)
	// c = z

}
