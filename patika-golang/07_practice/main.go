/*
1. studentName --> John Doe, grade --> 77, isPassed --> true değişkenlerini 3 farklı yöntem ile oluşturup, çıktısını yazdırınız.
2. Yukarıda belirtilen değişkenleri tek satır içerisinde tanımlayınız.
3. "Declaration", "Assign", "Initialization", "Initial Value" kavramlarını açıklayınız. (Terminoloji)
4. "Statically Typed" vs "Dynamically Typed" ifadelerini GO ve Python üzerinden gösteriniz.const
5. ":=" vs "=" aradaki farkı gösteriniz, double declaration
*/

// Package clause
package main

import "fmt"

func main() {

	// Cevap 1

	// var studentName string = "John Doe"
	// var grade int = 77
	// var isPassed bool = true

	// var studentName = "John Doe"
	// var grade = 77
	// var isPassed = true

	// studentName := "John Doe"
	// grade := 77
	// isPassed := true

	// var (
	// 	studentName string = "John Doe"
	// 	grade       int    = 77
	// 	isPassed    bool   = true
	// )

	// Cevap 2

	// var studentName, grade, isPassed = "John Doe", 77, true
	// studentName, grade, isPassed := "John Doe", 77, true

	// Cevap 3

	// var studentName string // "Declaration"
	// studentName = "John Doe" // "Assign"
	// var studentName string = "John Doe" // "Initialization"

	// Cevap 4

	// Go

	// Cevap 5

	// var studentName string = "John Doe"
	// studentName = "Aaa Bbb"

	studentName := "John Doe"
	studentName = "Aaa Bbb"

	fmt.Println(studentName)
	// fmt.Println(grade)
	// fmt.Println(isPassed)

}
