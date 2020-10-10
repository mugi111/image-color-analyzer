package image

import (
	"image"
)

func getPixelColor(img image.Image, x, y int) (uint32, uint32, uint32) {
	r, g, b, _ := img.At(x, y).RGBA()
	return r, g, b
}
