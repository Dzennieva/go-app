package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func projectPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/projects.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

func main() {
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/projects", projectPage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	log.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
