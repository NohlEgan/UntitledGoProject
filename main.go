package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func tenHelloWorldsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/tenhelloworlds" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, strconv.Itoa(i+1)+": Hello World!\n")
	}
}

func rightWayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	if r.URL.Path == "/therightway" {
		fmt.Fprintf(w, "You got to this page the right way!")
	} else if r.URL.Path == "/thewrongway" {
		fmt.Fprintf(w, "You got to this page the wrong way!")
	} else {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/tenhelloworlds", tenHelloWorldsHandler)
	http.HandleFunc("/therightway", rightWayHandler)
	http.HandleFunc("/thewrongway", rightWayHandler)

	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
