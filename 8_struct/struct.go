package main

import "fmt"

type Person struct {
	name   string
	age    int
	height float32
}

func (p Person) Greet(name string, age int, height float32) {
	p.name = name
	p.age = age
	p.height = height

	fmt.Println("call in function:-", p)

}

// func greet(p Person) {
// 	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
// }

func main() {

	var noman Person

	fmt.Println("Noman=>", noman.name)
	noman.Greet("Noman", 45, 5.6)
	fmt.Println("after=>", noman.name)

	// showVehicle()

}

// type Vehicle struct {
// 	wheel        int
// 	brand        string
// 	loadCapacity float32
// }

// type Truck struct {
// 	name string
// 	Vehicle
// }

// func showVehicle() {
// 	truck := Truck{
// 		name: "Lori Trunc BB4",
// 		Vehicle: Vehicle{
// 			wheel:        10,
// 			brand:        "Mercedes Benz",
// 			loadCapacity: 600.60,
// 		},
// 	}

// 	fmt.Println("Show New Vehicle")
// 	fmt.Println(truck)

// }
