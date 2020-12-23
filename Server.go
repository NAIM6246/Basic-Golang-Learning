package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	port := ":3000"
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/get", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome\n"))
		fmt.Fprintf(w, "Data Fetching")
	})

	r.Post("/post", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Your data is posted"))
	})

	r.Put("/put", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Data updated")
	})

	r.Delete("/delete", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Data is deleted")
	})

	fmt.Println("Serving on port ", port)
	http.ListenAndServe(port, r)
}
