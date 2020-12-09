package main

import (
  "fmt"
	"net/http"
)

// Handler main handler
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>image-color-analyzer</h1>")
}
