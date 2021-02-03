package main

import "fmt"

type person struct {
	name   string
	age    int
	Gender string
}

func main() {

	p1 := person{name: "Yash", age: 21}
	p2 := person{name: "Kaka", age: 1000}
	p3 := person{name: "Baba", age: 1500}

	fmt.Println("P1 : ", p1.name, p1.age)

	fmt.Println("P2 : ", p2.name, p2.age)

	fmt.Println("P3 : ", p3.name, p3.age)

}
