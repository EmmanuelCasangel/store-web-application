package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name, Description string
	Price             float64
	Quantity          int
}

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {

	products := []Product{
		{"Relogio", "Reluzente", 400, 1},
		{"Bolsa", "Carrega coisas", 40, 100},
	}

	tmpl.ExecuteTemplate(w, "Index", products)

}
