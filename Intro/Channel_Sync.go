package main

import (
	"fmt"
	"time"
)

func start(Done chan bool) {

	fmt.Print("Starting................")
	time.Sleep(time.Second)
	fmt.Print("Yal Margib")
}

func hellow() {

	fmt.Print("Hello......World")

}

func main() {

	done := make(chan bool, 1)

	go start(done)

	<-done

	go hellow()

}
