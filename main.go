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
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "SOmething went wrong", http.StatusBadGateway)
			return
		}

		fmt.Fprintf(w, "Relaying message : %s", b)
		w.WriteHeader(http.StatusOK)
		log.Println("Hello World")
	})

	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Bye bye devs")
	})

	http.ListenAndServe(":8000", nil)
}
