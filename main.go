package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/list", handler)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from Go.")
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid Method")
		return
	}
}
