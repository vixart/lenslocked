package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/sixsat/lenslocked/views"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "contact.html"))
}

func executeTemplate(w http.ResponseWriter, filePath string) {
	tpl, err := views.Parse(filePath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "home.html"))
}
