// Package clause
package main

import "fmt"

func main() {
	var name, surname string = "Kartal", "Kalyoncu"

	place := "Ankara"

	var age int = 35
	var isMarried bool = true

	place = "Yalova"

	var (
		name2      string = "Deneme"
		age2       int    = 33
		isMarried2 bool   = false
	)

	var name3, age3, isMarried3, weight = "aaa", 3, true, 72.5

	// name, age, isMarried, weight, height := "aa", 4, false, 72.5, 172

	var name4 string
	var age4 int

	fmt.Println("Hello World!")
	fmt.Println("Name Surname :", name, surname)
	fmt.Println("Name2 :", name2)
	fmt.Println("Name3 :", name3)
	fmt.Println("Place :", place)
	//fmt.Println("Surname :", surname)
	fmt.Println("Age :", age)
	fmt.Println("Age :", age2)
	fmt.Println("Age :", age3)
	fmt.Println("Weight :", weight)
	fmt.Println("Is Married :", isMarried)
	fmt.Println("Is Married :", isMarried2)
	fmt.Println("Is Married :", isMarried3)

	fmt.Println("Name4 :", name4)
	fmt.Println("Age :", age4)

}
