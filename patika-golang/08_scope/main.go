package main

import "fmt"

var packVar = "Package Scope"
var funcVar = "Func(Package) Scope"

func main() {

	if true {
		var ifVar = "If Scope"
		fmt.Println(ifVar)
	}

	var funcVar = "Func Scope"
	//funcVar := "Func Scope"

	fmt.Println(funcVar)

	fmt.Println(packVar)

	myFunc()

	var name = "Deneme"
	// name = "Yeni Deneme"
	// name := "Hatalı Deneme"
	name, surname := "Denem2", "Surname"

	fmt.Println(name, surname)
}

func myFunc() {
	fmt.Println(funcVar)
}
