package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloWorld)

	log.Println("Listening...")

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
