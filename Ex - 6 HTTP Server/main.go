package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	cts := req.Context()
	fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

	select{
		case <- time.After(5 * time.Second)
			fmt.Println(w,"hello")
		
	case <- cts.Done():

		err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
	}


}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
