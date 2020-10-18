package main

import (
	"image"
	"image/color"
)

type Image struct{ Width, Height int }

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x * y), 0, 255, 255}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
