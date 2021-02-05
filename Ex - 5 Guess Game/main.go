package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rNum(max int) int {
	
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(max)
	return rand.Intn(max)

}

func main() {

	fmt.Println("This is a Guessing game")
	var in int
	count := 0
	tries := 10

	num := rNum(100)

	userinp := make(chan int)

	fmt.Print("\nTake a Guess (You only get ", tries, " tries and max number is 100): ")

	for count < tries {

		go func() {
			fmt.Scanf("%d\n", &in)
			userinp <- in
		}()

		count++

		select {

		case user := <-userinp:
			{
				if user > num {
					fmt.Print("Guess Lower : ")
				} else if user < num {
					fmt.Print("Guess Higher : ")
				} else {
					fmt.Println("\nCorrect !!!!")
					fmt.Println("You Got the number right in ", count+1, "tries")
					return
				}
			}
		}

	}

	if count == tries {
		fmt.Println("\nYou Have Lost this Round \n The number was ", num)
	}

}
