package main

import (
	"fmt"
	"math/rand"
)

func main() {

	a := make([]int, 1)
	a[0] = 0
	i := 1

	for ; i <= 10; i++ {
		a = append(a, rand.Intn(100))

	}

	fmt.Println(a, " : ", len(a))

	s1 := a[2:5]

	fmt.Print(s1, " : ", len(s1))

}
