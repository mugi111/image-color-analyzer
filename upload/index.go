package upload

import (
	"net/http"
	"fmt"
)

// Handler upload handler
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid Method")
	}

	file, header, err := r.FormFile("file")

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fileHeader := make([]byte, 510)

	if _, err := file.Read(fileHeader); err != nil {
		fmt.Fprintf(w, "err read")
		return
	}

	file.Seek(-1, 0)

	defer file.Close()

	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintf(w, "File uploaded successfully: ")
	fmt.Fprintf(w, header.Filename)
	fmt.Fprintf(w, "MIME: %#v\n", http.DetectContentType(fileHeader))
}
