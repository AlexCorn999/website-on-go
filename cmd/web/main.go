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
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("starting app on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
