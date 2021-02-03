package main

import (
	"fmt"
	"time"
)

func main() {

	timer_1 := time.NewTimer(2 * time.Second)

	<-timer_1.C
	fmt.Println("Timer 1 has been fired")

	timer_2 := time.NewTimer(time.Second)

	go func() {
		<-timer_2.C
		fmt.Println("Timer 2 has been fired")
	}()

	stop2 := timer_2.Stop()

	if stop2 {
		fmt.Println("Timer Stopped")
	}

}
