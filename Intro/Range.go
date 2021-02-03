package main

import "fmt"

func main() {

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0

	for i, ele := range list {
		sum = sum + ele
		fmt.Println(i, " : ", sum)
	}

}
