package main

import "fmt"

func main() {
	fmt.Println("Start Main")
	fooo()
	fmt.Println("End Main")
}

func fooo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover From Panic")
		}
	}()

	fmt.Println("Start Foo")
	panic("Something went wrong")
	fmt.Println("End Foo")
}
