package main

import (
	"emmanuel/store-web-application/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8050", nil)
}
