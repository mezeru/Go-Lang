package main

import "fmt"

func main() {
	i := 0

	for i < 5 {
		fmt.Print(i, "\t")
		i++
	}
	fmt.Print("\n")

	for i := 10; i > 0; i-- {
		fmt.Print(i, "\t")
	}

}
