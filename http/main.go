package main

import (
	"net/http"
	"fmt"
	"os"
)

func main(){
	resp,err := http.Get("http://google.com")

	if err != nill {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}