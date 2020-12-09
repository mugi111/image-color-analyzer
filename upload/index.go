package upload

import (
	"image"
	"net/http"
	"fmt"
	"io"
)

func getPixels(file io.Reader) (x, y int, err error) {
  img, _, err := image.Decode(file)

  if err != nil {
    return 0, 0, err
  }

  bounds := img.Bounds()
	return bounds.Max.X, bounds.Max.Y, nil
}

func getPixelColor(img image.Image, x, y int) (uint32, uint32, uint32) {
	r, g, b, _ := img.At(x, y).RGBA()
	return r, g, b
}

// Handler upload handler
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid Method")
	}

	file, header, err := r.FormFile("file")

	if err != nil {
		fmt.Fprintln(w, "error:", err)
		return
	}

	fileHeader := make([]byte, 510)

	if _, err := file.Read(fileHeader); err != nil {
		fmt.Fprintf(w, "err read")
		return
	}

	file.Seek(-1, 0)

	width, height, err := getPixels(file)
	if err != nil {
		fmt.Fprintln(w, "error:", err)
	} else {
		fmt.Fprintln(w, width, height)
	}

	defer file.Close()

	fmt.Fprintf(w, "File uploaded successfully: ")
	fmt.Fprintf(w, header.Filename)
	fmt.Fprintf(w, "MIME: %#v\n", http.DetectContentType(fileHeader))
}
