package main

import "fmt"

func main() {

	msg := make(chan int)

	go func(a, b int) {
		msg <- a + b
	}(5, 2)

	fmt.Print(<-msg)

}
