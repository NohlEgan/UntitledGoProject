package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "./static/index.html")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		errorHandler(w, r, http.StatusNotFound)
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
		errorHandler(w, r, http.StatusNotFound)
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
		errorHandler(w, r, http.StatusNotFound)
	}
}

func fruitHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/fruit" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "./static/fruit.html")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	//w.WriteHeader(status)
	if status == http.StatusNotFound {
		http.ServeFile(w, r, "./static/404.html")
	}
}

func main() {
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/tenhelloworlds", tenHelloWorldsHandler)
	http.HandleFunc("/therightway", rightWayHandler)
	http.HandleFunc("/thewrongway", rightWayHandler)
	http.HandleFunc("/fruit", fruitHandler)

	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
