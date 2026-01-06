package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "base.html", map[string]string{"Title": "Home"})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "base.html", map[string]string{"Title": "About"})
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "base.html", map[string]string{"Title": "Order"})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/order", orderHandler)

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

/*package main

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

	//order route

	mux.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "order.html", nil)
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
