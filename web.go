package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"os"
	"html/template"
	"io"
	"strconv"
	"time"
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
	t, _ := template.ParseFiles("client.html")
        t.Execute(res, nil)
	fmt.Fprintln(res, "hello, heroku")
}



 

