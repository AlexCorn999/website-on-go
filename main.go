package main

import (
	"fmt"
	"log"
	"net/http"
)

var portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("starting app on port %s\n", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, nil))
}
