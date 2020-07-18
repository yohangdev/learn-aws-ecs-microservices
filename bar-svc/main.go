package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", BarServer)
	http.ListenAndServe(":8080", nil)
}

func BarServer(w http.ResponseWriter, r *http.Request) {
	hostName, _ := os.Hostname()

	log.Print("Logging in Bar! From: " + hostName)
	fmt.Fprintf(w, "Hello, Bar! From: " + hostName)
}