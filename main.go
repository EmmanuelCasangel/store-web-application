package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Product struct {
	id                int
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

	db := connectToDB()

	selectProducts, err := db.Query("select * from products")

	p := Product{}
	products := []Product{}

	for selectProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err)
		}
		p.id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
		products = append(products, p)
	}

	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "Index", products)

	defer db.Close()
}

func connectToDB() *sql.DB {
	conection := "user=postgres dbname=postgres host=localhost password=[PASS_WORD] sslmode=disable"

	db, err := sql.Open("postgres", conection)

	if err != nil {
		panic(err)
	}

	return db

}
