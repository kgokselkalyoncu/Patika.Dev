package main

import "fmt"

func main() {

	// x := 10
	// y := 10.0

	// fmt.Printf("%v %T\n", x, x)
	// fmt.Printf("%v %T\n", y, y)

	// Type Conversion type(value) => int(y)

	// fmt.Println(x + int(y))

	/*
		var x int8 = 10
		var y int16 = 10

		fmt.Println(x + y)
	*/

	/*
		x := 10
		y := 10.0

		fmt.Printf("%v %T\n", x, x)
		fmt.Printf("%v %T\n", y, y)

		fmt.Println(float64(x) + y)
	*/

	/*
		var x int16 = 128
		var y int8

		y = int8(x)

		fmt.Println(y)
	*/

	/*
		x := 10
		y := "10"

		fmt.Printf("%v %T\n", x, x)
		fmt.Printf("%v %T\n", y, y)

		fmt.Printf(x + int(y))
	*/

	num1 := 106
	str1 := string(rune(num1))

	fmt.Printf("%v %T\n", num1, num1)

	fmt.Printf("%v %T\n", str1, str1)

}
