package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Creating server instance in port 8000")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
	})

	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Bye bye devs")
	})

	http.ListenAndServe(":8000", nil)
}
