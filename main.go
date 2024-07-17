package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	fmt.Fprintf(w, "Hello World! %s", time.Now())
	response := []byte("hello world")
	_, err := w.Write(response)
	if err != nil {
		log.Panic("error while writing the response")
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	w.Write([]byte("hello, I am alive..."))
	log.Println("hello, I am alive...")
}
func main() {
	http.HandleFunc("/hello-world", greet)
	http.HandleFunc("/health", handleHealth)
	address := "localhost:8000"
	log.Printf("Listening on %s", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
