package main

import "fmt"

var p = 10   // run time execute
const a = 20 // compile time execute

func outer() func() {
	age := 24
	networth := 100

	fmt.Printf("Noman is %v years old\n", age)

	var calculateNetWorth = func() {
		networth = networth * 2

		fmt.Printf("Noman Networth is  %v\n", networth)
	}

	return calculateNetWorth
}

func call() {

	outer1 := outer()
	outer1()
	outer1()

	outer2 := outer()
	outer2()
	outer2()
}

func main() {
	call()
}

/*

MY EXPECTED OUTPUT

	Noman is 24 years old

	Noman Networth is 200
	Noman Networth is 400

	Noman is 24 years old

	Noman Networth is 200
	Noman Networth is 400

	=======================


Phases:
  1. compilation phase (compile time)
  2. execution phase (execution time)

  ========= COMPILE PHASE ===========

  * code segment *


  * data segment (global memory) *


  * stack *


  * heap (GC => garbage collector) *


*/
