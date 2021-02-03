package main

import (
	"fmt"
	"time"
)

func main() {

	tincker := time.NewTicker(100 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return

			case t := <-tincker.C:
				fmt.Println("Thinker at : ", t)

			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	tincker.Stop()
	done <- true

	fmt.Println("Ticker Stopped")

}
