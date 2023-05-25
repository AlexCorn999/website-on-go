package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlexCorn999/website-on-go/pkg/config"
	"github.com/AlexCorn999/website-on-go/pkg/handlers"
	"github.com/AlexCorn999/website-on-go/pkg/render"
)

var portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("starting app on port %s\n", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, nil))
}
