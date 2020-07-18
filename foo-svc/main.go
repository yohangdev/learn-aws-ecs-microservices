package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", FooServer)
	http.ListenAndServe(":8080", nil)
}

func FooServer(w http.ResponseWriter, r *http.Request) {
	hostName, _ := os.Hostname()

	log.Print("Logging in Foo! From: " + hostName)
	fmt.Fprintf(w, "Hello, Foo! From: " + hostName)
}