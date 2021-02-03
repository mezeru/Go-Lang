package main

import "fmt"

func main() {

	channels := make(chan int, 9)

	channels <- 1
	channels <- 8
	channels <- 1
	channels <- 2
	channels <- 1
	channels <- 0
	channels <- 0
	channels <- 6
	channels <- 2

	for i := 0; i < 9; i++ {

		fmt.Print(<-channels)

	}

}
