package main

import "fmt"

func (a, b int) int {

	return a + b

}

func main() {

	var a, b int

	fmt.Println("Enter the value of a and b : ")
	fmt.Scanln(&a, &b)
	ans := sum(a, b)

	fmt.Println("The Sum is : ", ans)

}
