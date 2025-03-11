package main

import "fmt"

func main() {
	// Arethmetic Operators

	// var a int = 10
	// var b int = 20

	// fmt.Println("Addition:", a+b)
	// fmt.Println("Subtraction:", a-b)
	// fmt.Println("Multiplication:", a*b)
	// fmt.Println("Division:", b/a)
	// fmt.Println("Modulus:", a%b)

	// Logical Operators

	// x := true
	// y := false

	// fmt.Println("Logical AND:-", x && y)
	// fmt.Println("Logical OR:-", x || y) // IF GET 1 TRUE THEN RETURN TRUE or RETURN FALSE
	// fmt.Println("Logical NOT:-", !x)

	// Comparison Operators

	// x, y := 10, 20

	// fmt.Println("Equal:", x == y)     // false
	// fmt.Println("Not Equal:", x != y) // true
	// fmt.Println("Greater:", x > y)    // false
	// fmt.Println("Less:", x < y)       // true

	// Conditional Statements

	// num := 10

	// if num > 20 {
	// 	fmt.Println("Number is gretter than 20")
	// } else if num < 20 {
	// 	fmt.Println("Number is small then 20")
	// } else {
	// 	fmt.Printf("Exact Number is %v", num)
	// }

	// Switch Case

	day := "Noman2"

	switch day {

	case "Noman":
		fmt.Println("This is Noman")
	case "Ali":
		fmt.Println("This is Ali")
	case "Khan":
		fmt.Println("This is Khan")

	default:
		fmt.Println("Invalid Person")
	}
}
