package main

import (
	"fmt"
	"strings"
)

func main() {

	count := 1

	var input string
	fmt.Scanf("%s\n", &input)
	fmt.Println("Input is ", input)

	for _, char := range input {

		str := string(char)

		if strings.ToUpper(str) == str {
			count++
		}
	}

	fmt.Println("There are ", count, " words in the string ")

}
