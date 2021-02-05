package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, i := strconv.ParseInt("234324", 0, 64)
	fmt.Println(f, i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	e, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(e)

	_, err := strconv.Atoi("wat")
	fmt.Println(err)

}
