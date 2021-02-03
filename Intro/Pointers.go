package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)

	var ptr *int
	ptr = &a

	fmt.Print("A : ", a, "\nptr : ", ptr, "\n*ptr : ", *ptr)

}
