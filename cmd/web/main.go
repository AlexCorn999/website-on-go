package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AlexCorn999/website-on-go/pkg/config"
	"github.com/AlexCorn999/website-on-go/pkg/handlers"
	"github.com/AlexCorn999/website-on-go/pkg/render"
	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var portNumber = ":8080"
var session *scs.SessionManager

func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
