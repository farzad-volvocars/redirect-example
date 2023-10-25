package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func handleContinue(w http.ResponseWriter, r *http.Request) {
	log.Print("handleContinue called")

	appId := r.URL.Query().Get("id")
	http.Redirect(w, r, fmt.Sprintf("http://localhost:3000/signed?id=%v", appId), http.StatusSeeOther)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/continue", handleContinue)
	http.ListenAndServe(":3003", r)
}
