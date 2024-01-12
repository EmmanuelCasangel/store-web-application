package models

import (
	"emmanuel/store-web-application/db"
	"log"

	"github.com/google/uuid"
)

type Product struct {
	Id                string
	Name, Description string
	Price             float64
	Quantity          int
}

func GetAllProducts() []Product {

	db := db.ConnectToDB()
	selectProducts, err := db.Query("select * from store_products")

	p := Product{}
	products := []Product{}

	for selectProducts.Next() {
		var id string
		var quantity int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err)
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
		products = append(products, p)
	}

	if err != nil {
		panic(err)
	}
	defer db.Close()

	return products
}

func CreateProduct(name, descripiton string, price float64, quantity int) {
	db := db.ConnectToDB()

	insertData, err := db.Prepare("insert into store_products values($1, $2, $3, $4, $5)")
	if err != nil {
		panic(err.Error())
	}

	id := uuid.NewString()

	log.Println("id", id, "nome", name, "des", descripiton, "pr", price, "quantity", quantity)

	insertData.Exec(id, name, descripiton, price, quantity)
	defer db.Close()
}

func RemoveProduct(id string) {
	db := db.ConnectToDB()

	removeProduct, err := db.Prepare("delete from store_products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	removeProduct.Exec(id)
	defer db.Close()
}
