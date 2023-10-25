package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	log.Print("handleHomePage called")

	t, err := template.ParseFiles("index.html")
	if err != nil {
		w.Write([]byte("Client Error"))
	}

	t.Execute(w, nil)
}

func handleSigned(w http.ResponseWriter, r *http.Request) {
	log.Print("handleSigned called")

	w.Write([]byte("Now signed"))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handleHomePage)
	r.Get("/signed", handleSigned)
	http.ListenAndServe(":3000", r)
}
