package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /webhook", func(w http.ResponseWriter, r *http.Request) {
		// print the request body
		fmt.Println(r.Body)
	})

	fmt.Println("Server is running on port 5050")
	http.ListenAndServe(":5050", mux)
}
