package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Using var

	// var name string = "Noman"
	// var age int = 23
	// var isCool bool = true
	// var size float32 = 2.3

	// fmt.Printf("Name: %v\n, Age: %v\n, isCool: %v\n, size: %v\n", name, age, isCool, size)

	// Using shorthand

	// name1 := "Noman1"
	// age1 := 23
	// isCool1 := true

	// fmt.Printf("Name: %v\n, Age: %v\n, isCool: %v\n", name1, age1, isCool1)

	// Using const

	// const a int = 10
	// const b = 20

	// const c = a + b

	// fmt.Println("sum a and b = ", c)

	// Using multiple variables

	// var name, email string = "Noman", "noman@gmail"

	// fmt.Println(name, email)

	// Using multiple variables with shorthand

	// name1, email1 := "Noman1", "noman1@gmail"

	// fmt.Println(name1, email1)

	// using multiple variables with different types (Grouped)

	// var (
	// 	name2   string = "Noman2"
	// 	age2    int    = 23
	// 	isCool2 bool   = true
	// )

	// name2 = "Noman3 reassign"

	// println(name2, age2, isCool2)

	// String and characters in Go lang

	// var character byte = 'A'

	// var sybmol rune = 'A'

	// println(character, sybmol)

	// Type Conversion in Go lang

	// intiger to float conversion
	// var x int = 10
	// var intToFloat float64 = float64(x)
	// println(intToFloat)

	// float to int conversion
	// var y float32 = 10.533
	// var floatToInt int = int(y)
	// println(floatToInt)

	// int to string conversion
	var z int = 10
	var intToString string = strconv.Itoa(z)
	fmt.Println(intToString)

	// string to int conversion
	var str string = "K"

	stringToInt, _ := strconv.Atoi(str)
	// if err != nil {
	// 	fmt.Println("Error converting string to int:", err)
	// } else {
	// 	fmt.Println(stringToInt)
	// }
	fmt.Println(stringToInt)
}
