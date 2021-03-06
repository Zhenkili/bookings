package main

import (
	"net/http"

	"github.com/Zhenkili/bookings/pkg/config"
	"github.com/Zhenkili/bookings/pkg/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//use router instead of main
func routes(app *config.Appconfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(Nosurf)
	mux.Use(SessionLoad)
	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)

	//connect the server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
