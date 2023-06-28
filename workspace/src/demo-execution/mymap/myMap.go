package mymap

import "fmt"

func MyMap() {

	// Go provide 2 ways to declare a map
	// each ways have the same display value but zero value is different
	// syntax to define map type: map[<key type>]<value type>

	var myMap1 map[string]int

	fmt.Println(myMap1)
	if myMap1 == nil {
		fmt.Println("myMap1 is nil.")
	} else {
		fmt.Println("myMap1 is not nil.")
	}

	var myMap2 = make(map[string]int)

	fmt.Println(myMap2)

	if myMap2 == nil {
		fmt.Println("myMap2 is nil.")
	} else {
		fmt.Println("myMap2 is not nil.")
	}

	myMap3 := map[string]int{
		"top":    1,
		"bottom": 2,
	}

	fmt.Println(myMap3)

	// add element into map
	myMap3["left"] = 6
	fmt.Println(myMap3)

	// delete element in map
	delete(myMap3, "top")
	fmt.Println(myMap3)
	fmt.Println(len(myMap3))

	// map is also a reference type
	myMap4 := myMap3
	myMap4["right"] = 5

	fmt.Println(myMap3)

	// when access value in map it will return 2 value
	// 1. value of key // will be zero value if key not in map
	// 2, A value check if the key is in map or not (bool)

	value, found := myMap4["top"]

	if found {
		fmt.Println(value)
	} else {
		fmt.Println("Not found!")
	}

	// Go cannot compare 2 map: myMap1 == myMap2
}
