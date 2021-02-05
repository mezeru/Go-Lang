package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {

	response, err := http.Get("https://go.dev/")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	mydir, err := os.Getwd()
	mydir = mydir + "/index.html"
	f, err := os.Create(mydir)

	fmt.Println("Response Status : ", response.Status)

	scanner := bufio.NewScanner(response.Body)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		f.WriteString(scanner.Text())
	}

}
