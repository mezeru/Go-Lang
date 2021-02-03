package main

import "fmt"

func Factorial(n int) int {
	3

	if n == 1 {
		return 1
	}

	return n * Factorial(n-1)

}

func main() {

	var in int

	fmt.Print("Enter a number : ")
	fmt.Scan(&in)

	fmt.Print("The Factorial of ", in, " is : ", Factorial(in))
}
