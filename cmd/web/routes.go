package main

import (
	"net/http"

	// "github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mitanshukr/bookings-app/internal/config"
	"github.com/mitanshukr/bookings-app/internal/handlers"
)

func Routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	// todo: session middlewares

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Get("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	

	fileServer := http.FileServer(http.Dir("./static/"))

	// http.StripPrefix("/static", fileServer) creates a handler that removes the /static prefix from incoming requests before passing them to fileServer.
	// This is necessary because fileServer expects requests without the /static prefix.
	// So, any request that comes in as /static/file.html will be translated to file.html before being handled by fileServer.
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
