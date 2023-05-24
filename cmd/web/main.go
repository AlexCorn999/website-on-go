package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlexCorn999/website-on-go/pkg/handlers"
)

var portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("starting app on port %s\n", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, nil))
}
