package main

import "fmt"

func main() {

	var ar [5]int

	for i := 0; i < 5; i++ {

		fmt.Scanln(&ar[i])

	}

	for i := 0; i < 5; i++ {

		fmt.Println("i[", i, "] = ", ar[i])

	}

}
