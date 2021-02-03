package main

import "fmt"

func main() {

	a := 181210000

	sum := 0

	for true {

		a++

		if a > 181210066 {
			break
		}

		sum = sum + a

	}

	fmt.Print("The mean is : ", sum/66)
}
