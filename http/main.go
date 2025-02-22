package main

import (
	"net/http"

	"example.com/http/routes"
)

func main() {
	r := routes.MainRouter

	routes.BooksApi()

	http.ListenAndServe(":8080", r)
}
