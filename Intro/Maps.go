package main

import "fmt"

func main() {

	l1 := make(map[string]int)

	names := []string{"Yash", "Ruizolas",
		"Mezeru",
		"Jermpro0",
		"Purple"}

	for i := 0; i < len(names); i++ {

		l1[names[i]] = len(names[i])*10 + 21 // Hash Function

	}

	fmt.Println(l1)

	delete(l1, "Yash")
	fmt.Println(l1)

}
