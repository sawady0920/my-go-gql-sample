package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/aaa", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "aaa\n")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "AAA\n")
	})

	http.ListenAndServe(":8080", nil)
}
