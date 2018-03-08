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
	
			}
}

func hello(res http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadFile("domains.txt") // just pass the file name
    	if err != nil {
        fmt.Print(err)
			}
	str := string(b)
	fmt.Fprintln(res, str) 
}
	

