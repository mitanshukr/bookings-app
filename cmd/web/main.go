package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mitanshukr/bookings-app/internal/config"
	"github.com/mitanshukr/bookings-app/internal/handlers"
	"github.com/mitanshukr/bookings-app/internal/render"
)

const portNumber = ":8080"

func main() {
	var appConfig config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Panicln("unable to create template cache", err)
	}

	appConfig.TemplateCache = tc
	appConfig.IsDev = true

	render.NewTemplate(&appConfig)

	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("starting go server on port", portNumber)
	// http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&appConfig),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
