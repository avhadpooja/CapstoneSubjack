package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	b, err := ioutil.ReadFile("file.txt") // just pass the file name
    	if err != nil {
        fmt.Print(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, b)
}

