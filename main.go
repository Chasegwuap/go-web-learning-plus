package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Serve static files (CSS)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	// Load HTML template
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Home route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	log.Println("Server running🚀🚀 on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
