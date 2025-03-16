package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divided(10, 3)

	if err != nil {
		fmt.Println("Errors", err)
	} else {
		fmt.Println("Result", result)
	}

}

func divided(a float32, b float32) (float32, error) {

	if b == 0 {
		return b, errors.New("division by zero is not allowed")
	} else {
		return a / b, nil
	}
}
