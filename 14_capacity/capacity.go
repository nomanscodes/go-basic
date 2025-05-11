package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// To avoid memory issues, consider using a smaller number like 1000000
	items := make([]int, 0, 100000000) // Preallocate capacity

	for i := 0; i < 10000000000; i++ {
		items = append(items, i)
	}

	elapsed := time.Since(start)


	fmt.Println("Execution time =>", elapsed)
}
