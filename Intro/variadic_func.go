package main

import "fmt"

func vari(sli ...int) {
	print(sli)
}

func main() {

	sli := make([]int, 1)
	var a, in int
	fmt.Print("Enter the size of the list : ")
	fmt.Scanln(&a)

	fmt.Println("Enetr the elements : ")
	for i := 0; i < a; i++ {
		fmt.Scanln(&in)
		append(sli, in)
	}

	vari(sli...)

}
