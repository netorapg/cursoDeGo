package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	u := usuario{
		"Renato Augusto",
		"reh@gmail.com",
	}
	templates.ExecuteTemplate(w, "home.html", u)
}

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	fmt.Println("Server is running at port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
