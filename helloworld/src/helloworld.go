package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello 1.0.0!")
	})

	http.ListenAndServe(":8080", nil)
}
