package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 127, 127)
}

func (i Image) At(x, y int) color.Color {
	R := (x*x + y*y)/30
	G := 255
	B := (x*x + y*y)/10
	A := 255
	return color.RGBA{uint8(R), uint8(G), uint8(B), uint8(A)}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
