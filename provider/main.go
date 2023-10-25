package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func handleLoginForm(w http.ResponseWriter, r *http.Request) {
	log.Print("handleLoginForm called")

	appId := r.URL.Query().Get("id")

	t, err := template.ParseFiles("index.html")
	if err != nil {
		w.Write([]byte("Auth Error"))
	}

	t.Execute(w, appId)

}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/auth", handleLoginForm)
	http.ListenAndServe(":3001", r)
}
