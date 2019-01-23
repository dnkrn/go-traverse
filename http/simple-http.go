package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", simple)
	http.HandleFunc("/file", fileServe)
	startError := http.ListenAndServe(":8080", nil)
	if startError != nil{
		fmt.Print("error starting application")
	}
	
}

func simple(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "My name is Dinakaran.")
}

func fileServe(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "/Users/dinakaran/Documents/Interview-Hints.pages")
	http.ServeFile(res, req, "/Users/dinakaran/Documents/Interview-Hints.pages")
}