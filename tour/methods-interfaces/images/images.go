package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rectangle{image.Pt(0, 0), image.Pt(100, 100)}
}

func (i Image) At(x, y int) color.Color {
	 return color.RGBA{0 + uint8(x), 0 + uint8(y), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
