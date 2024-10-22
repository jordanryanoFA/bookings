package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/jordanryanoFA/bookings/internal/config"
	"github.com/jordanryanoFA/bookings/internal/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// //  RenderTemplate renders templates using html/templates (the complex version)
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app Config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

// createTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files name *.go.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.go.tmpl")
	if err != nil {
		return myCache, err
	}

	// ranges through all files ending with *go.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.go.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.go.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}

// // //  this is the beginner version

// // RenderTemplate renders templates using html/templates
// func RenderTemplate(w http.ResponseWriter, tmpl string) {

// 	parseTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.go.tmpl")
// 	err := parseTemplate.Execute(w, nil)

// 	if err != nil {
// 		fmt.Println(" error parsing template:", err)
// 		return
// 	}
// }
// var tc = make(map[string]*template.Template)

// func RenderTemplateTest (w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	//check to see if we already have templates in our cache
// 	_, inMap := tc[t]
// 	if !inMap {
// 		//need to create a template
// 		log.Println("creating template and adding cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		// we have the template in the cache
// 		log.Println("using cache template")
// 	}

// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.go.tmpl",
// 	}
// 	//parse the templates
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}
// 	// add templates to cache (map)
// 	tc[t] = tmpl

// 	return nil
// }
