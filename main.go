package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go Web Server")
	})

	port := 3000
	fmt.Printf("Server is listening on 3000")
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Printf("Error: ", err)
	}
}
