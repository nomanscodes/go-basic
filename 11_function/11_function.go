package main

import "fmt"

/*
Higher Order Function
 1. Get a function by parameter
 2. Return Function
 3. Both

First order function (Work With rules)

CallBack Function

Function Expression

First Order Item
*/

// Higer order funciton Like Get Function
func makeOperation(a int, b int, sum func(x int, y int)) {
	sum(a, b)
}

// Higher order Function Like Return Another Funciton
func call() func(x, y int) {
	return calculateMultyply
}

func calculateMultyply(p, q int) {
	fmt.Println("Multiply:-", p*q)
}

func doSum(x, y int) {
	sum := x + y
	fmt.Println("Sum:-", sum)
}

func main() {
	fmt.Println("Call Main Funciton")

	makeOperation(5, 5, doSum)

	var callOutput = call()

	callOutput(10, 10)

	// anonymous function
	func(a, b int) {
		sum := a + b
		fmt.Println("Sum Form anonymous function", sum)
	}(5, 5)

	// function expression
	var sumFunc = func(a, b int) {
		sum := a + b

		fmt.Println("Sum From anonymous function expressions", sum)
	}

	// Invoke anonymous function
	sumFunc(10, 10)
}

// init funciton
func init() {
	fmt.Println("Call Init Funciton")
}
