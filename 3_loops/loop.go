package main

import "fmt"

func main() {

	// For Loop
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i)
	// }

	// While Loop BY using For Loop

	// j := 0

	// for j <= 100 {
	// 	fmt.Println(j)
	// 	j++
	// }

	// Infinite For Loop with Break

	// k := 1

	// for {
	// 	k++
	// 	fmt.Println("Infinite Loop", k)

	// 	if k == 20 {
	// 		break
	// 	}

	// }

	// For Loop with slice

	// numbers := []int{10, 20, 30, 40, 50}

	// for i := 0; i <= len(numbers)-1; i++ {
	// 	fmt.Println(numbers[i])
	// }

	// For Loop with Range

	// numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	// person := map[string]string{
	// 	"name":  "Noman",
	// 	"email": "noman@example.com",
	// 	"city":  "Dhaka",
	// }

	// for index, value := range numbers {
	// 	fmt.Println("Index:", index, "Value:", value)
	// }

	// For Loop with Range and _ (underscore)

	// for key, value := range person {
	// 	fmt.Println("Key:", key, "Value:", value)
	// }

	// name := "Noman"

	// for index, char := range name {
	// 	fmt.Println("Index:", index, "Char:", string(char))
	// }

	// EXERCISES

	// Print all Even numbers from 1 to 50

	for i := 1; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Println("Even Number:-", i)
		}
	}

	// Print all odd numbers from 1 to 50

	for i := 0; i <= 50; i++ {
		if i%2 != 0 {
			fmt.Println("Odd Number:-", i)
		}
	}
}
