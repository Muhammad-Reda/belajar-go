package methods

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You've requested all books")
}

func GetBooksTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	fmt.Fprintf(w, "The book's title is %s", title)
}

func GetAllInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]
	fmt.Fprintf(w, "The read %s on page %s", title, page)
}
