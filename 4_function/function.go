package main

import "fmt"

func main() {
	// function call
	// result := add(10, 20)
	// fmt.Println("Addition:", result)

	// Named return value
	// fmt.Println("Area of Ractangle:", ractangleArea(10.64, 2.01))

	// Variadic Function
	// fmt.Println("Summation:", Summation(10, 20, 30, 40, 50, 60, 70, 80, 90, 100))

	// Defer Function
	// deferFunc()

	// Panic Function
	panicFunc()

}

// function parameter and return type with value

// func add(a int, b int) int {
// 	return a + b
// }

// Named return value
// func ractangleArea(lenght float64, width float64) (area float64) {
// 	area = lenght * width
// 	return
// }

// Variadic Function

// func Summation(numbers ...int) (sub int) {

// 	result := 0

// 	for _, value := range numbers {
// 		result += value
// 	}

// 	return result
// }

// Defer Function

// func deferFunc() {
// 	fmt.Println("Start")

// 	// This is call after the function all the code is executed
// 	defer fmt.Println("Middle")

// 	fmt.Println("End")
// }

// Panic function

func panicFunc() {
	fmt.Println("Start")

	panic("Something went wrong")

	// This is not printing because the panic function is stop the execution of the code
	fmt.Println("End")
}
