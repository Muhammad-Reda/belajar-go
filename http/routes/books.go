package routes

import (
	"example.com/http/methods"
)

func BooksApi() {
	bookRouter := MainRouter.PathPrefix("/books").Subrouter()

	bookRouter.HandleFunc("/", methods.GetAllBooks)

	bookRouter.HandleFunc("/{title}", methods.GetBooksTitle)

	bookRouter.HandleFunc("/{title}/page/{page}", methods.GetAllInfo)

}
