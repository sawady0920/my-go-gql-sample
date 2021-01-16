package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/aaa", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "aaa\n")
		if err != nil {
			panic(err)
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "AAA\n")
		if err != nil {
			panic(err)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
