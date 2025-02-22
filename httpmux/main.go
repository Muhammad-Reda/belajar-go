package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}

		// Set header content type
		w.Header().Set("Content-Type", "application/json")

		// Create response data
		response := map[string]string{
			"message": "Welcome to home page",
		}

		err := json.NewEncoder(w).Encode(response)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	router.HandleFunc("/user/", func(w http.ResponseWriter, req *http.Request) {
		id := req.URL.Path[len("/user/"):]

		if id == "" {
			res := map[string]string{
				"message": "Params can't blank",
			}
			http.NotFound(w, req)
			json.NewEncoder(w).Encode(res)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		response := map[string]string{
			"message": "You've accessed user path",
			"id":      id,
		}

		err := json.NewEncoder(w).Encode(response)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
		return
	}
}
