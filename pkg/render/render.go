package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/AlexCorn999/website-on-go/pkg/config"
)

var app *config.AppConfig

// New Templates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// Render template renders a template
func RenderTemplate(w http.ResponseWriter, html string) {
	// get the template cache from the app config
	tc := app.TemplateCache

	// get requested template from cache
	t, ok := tc[html]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html from /templates
	pages, err := filepath.Glob("../../templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("../../templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
