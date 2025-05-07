package main

import "fmt"

func main() {

	var intializeArray [4]string = [4]string{
		"Noman",
		"Safiq",
		"Rafiq",
		"Mehedi",
	}

	fmt.Println("Array Initialize:-", intializeArray)

	// declare array
	var arr [5]int

	// array item data insert
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	fmt.Println("Array sliceing:-", arr[2:3])

	// array item modify
	arr[1] = 0

	fmt.Println(arr)

	// check array length
	fmt.Println(len(arr))

	// individual item access
	fmt.Println(arr[4])

	// array print by loop

	for i := range len(intializeArray) {

		if intializeArray[i] == "Noman" {
			fmt.Println(intializeArray[i])
		}

	}

}
