package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Static files
	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	// Load templates
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	// Home route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	// About route
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "about.html", nil)
	})

	log.Println("Server running 🚀 on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

/*package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Static files
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	// Load all templates
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	// Home
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	// About
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "about.html", nil)
	})

	log.Println("Server running 🚀 on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}




/*package main

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

	// About route
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About Page"))
	})

	log.Println("Server running🚀🚀 on http://localhost:8080")
	// log.Println("Server running🚀🚀 on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
``
*/
