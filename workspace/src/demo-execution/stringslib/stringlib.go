package stringslib

import (
	"fmt"
	"strconv"
	"strings"
)

func Stringslib() {

	str := "hello world!"
	// len() return the length of string
	fmt.Println(len(str))

	// String in Go are slice of bytes
	// So you can get the substring by using [<start index>:<end index>]
	// by default start index will be 0 if not defined
	// by default end index will be the last or (len(string) - 1) if not defined
	fmt.Println(str[0:5])
	fmt.Println(str[:5])
	fmt.Println(str[0:])

	nstr := []byte{97, 98, 99}
	fmt.Printf("%s\n", nstr)

	fmt.Println(string(nstr))

	// strings library
	// common function help working with string

	fmt.Println(strings.ToLower("HELLO"))
	fmt.Println(strings.ToUpper("hello"))
	// check string start with substring
	fmt.Println(strings.HasPrefix("hello", "h"))
	// check string end with substring
	fmt.Println(strings.HasSuffix("hello", "o"))

	fmt.Println(strings.Compare("S", "S"))
	fmt.Println(strings.Replace("helloww", "w", "", 1))
	fmt.Println(strings.ReplaceAll("helloww", "w", ""))

	// convert int -> string
	x := 3
	y := strconv.Itoa(x)

	fmt.Printf("type = %T value = %v\n", y, y)

	// convert string -> int
	z, _ := strconv.Atoi(y)
	fmt.Printf("type = %T value = %v\n", z, z)

}
