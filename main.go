package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func Calculate(x int) int {
	return x + 2
}

func main() {
	fmt.Println("Go CI Pipeline Tutorial")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":4000", nil))

}
