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

	http.ListenAndServe(":5050", mux)
}
