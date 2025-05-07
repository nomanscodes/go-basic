package main

import "fmt"

func main() {

	// declare slice
	var mySlice []int

	// append data in slice
	mySlice = append(mySlice, 10)
	mySlice = append(mySlice, 20)
	mySlice = append(mySlice, 30)

	// sliceing slice
	fmt.Println(mySlice[1:2])

	// Initialize slice
	var myslice2 []string = []string{
		"one",
		"two",
		"three",
		"four",
	}

	fmt.Println(myslice2)

	// append item in initilize slice
	myslice2 = append(myslice2, "five")

	// access individual item from slice
	fmt.Println(myslice2[3])

	// initilize a array
	var myarry [3]int = [3]int{
		1, 2, 3,
	}

	//  slice from an array
	var myarraySlice []int = myarry[0:2]

	fmt.Println("slice from an array:-", myarraySlice)

}
