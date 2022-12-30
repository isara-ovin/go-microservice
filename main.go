package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("Creating server instance in port 8000")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		log.Println(string(b))

		fmt.Fprintf(w, "Relaying message : %s", b)
		log.Println("Hello World")
	})

	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Bye bye devs")
	})

	http.ListenAndServe(":8000", nil)
}
