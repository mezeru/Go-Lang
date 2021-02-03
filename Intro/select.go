package main

import (
	"fmt"
	"time"
)

func main() {

	ci := make(chan string, 1)
	cii := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ci <- "Hello this is one"
	}()

	go func() {
		time.Sleep(time.Second)
		cii <- "Hello this is two"
	}()

	for i := 0; i < 2; i++ { // For loop for both msgs to arrive

		select {
		case msg1 := <-ci:
			fmt.Println(msg1)

		case msg2 := <-cii:
			fmt.Println(msg2)

		}

	}

}
