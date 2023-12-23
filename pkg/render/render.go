package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/mitanshukr/bookings-app/pkg/config"
	"github.com/mitanshukr/bookings-app/pkg/models"
)

var appConfig *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	appConfig = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// create a template cache
	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//improvement: get template cache from appConfig
	var tc map[string]*template.Template
	if appConfig.IsDev {
		tc, _ = CreateTemplateCache()
	} else {
		tc = appConfig.TemplateCache
	}

	// get requested template from the cache
	t := tc[tmpl]
	if t == nil {
		log.Fatal("unable to get templates from cache...")
	}

	//----- optionally writing to buf, for fail safety
	buf := new(bytes.Buffer)
	err := t.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.pages.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all the files ending with *.pages.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		// log.Println(matches)
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
		// log.Println(myCache)
		// log.Println("created template cache, success!")
	}

	return myCache, nil
}

// func RenderTemplateOld(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error parsing template", err)
// 		return
// 	}
// }

// var tc = make(map[string]*template.Template)

// func RenderTemplateWithCacheSimple(w http.ResponseWriter, t string) {
// 	// log.Println(tc)
// 	var tmpl *template.Template
// 	var err error

// 	tmpl = tc[t]

// 	if tmpl == nil {
// 		// create template cache
// 		log.Println("creating temp cache and render...")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Fatal("error parsing template & cache")
// 		}
// 		tmpl = tc[t]
// 	}

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error parsing template", err)
// 		return
// 	}
// 	// log.Println("hello...")
// 	// log.Println(tc)
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl
// 	return nil
// }
