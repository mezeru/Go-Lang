package main

import (
	"fmt"
	"time"
)

func fun(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str, " ", i)
	}
}

func main() {

	go fun("go 1")

	go func(str0 string) {
		fmt.Println(str0)
	}("This Call")

	time.Sleep(time.Second)

	fmt.Println("Done")
}
