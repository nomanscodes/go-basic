package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	var bptr *int = &b

	workers(a, bptr)

	fmt.Println("A in main:-", a)
	fmt.Println("B in main:-", b)

}

func workers(a int, b *int) {

	*b = 100

	result := a + *b
	fmt.Println("result:-", result)
	fmt.Println("B in workers-", b)
	fmt.Println("A in workers:-", a)
}
