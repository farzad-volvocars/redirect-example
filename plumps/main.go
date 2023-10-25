package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func handleCallback(w http.ResponseWriter, r *http.Request) {
	log.Print("handleCallback called")

	appId := r.URL.Query().Get("id")
	http.Redirect(w, r, fmt.Sprintf("http://localhost:3003/continue?id=%v;", appId), http.StatusSeeOther)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/callback", handleCallback)
	http.ListenAndServe(":3002", r)
}
