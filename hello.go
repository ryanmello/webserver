package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method Not Supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "HELLO!")
}

func formHandler() {
	
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}