package main

import (
	"fmt"
)

func opertions(a, b int) (int, int, int, int) {

	return a + b, a - b, a * b, a / b

}

func main() {

	var a, b int

	fmt.Print("Enter the value of a and b : ")
	fmt.Scanln(&a, &b)
	sum, diff, mul, div := opertions(a, b)

	fmt.Println("The Sum is : ", sum, "\nDifference is :", diff, "\nThe Product is : ", mul, "\nThe Division is : ", div)

}
