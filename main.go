package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", helloHandler)

	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello")
}
