package main

import "fmt"

func main() {

	c1 := make(chan int, 2)
	c1 <- 34
	c1 <- 32
	close(c1)

	for i := range c1 {
		fmt.Println(i)
	}

}
