package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	"math/rand"
	"time"
)

type Image struct{
	Width, Height int
	colours [][]color.RGBA

}

func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i *Image) At(x, y int) color.Color {
	return i.colours[x][y]
}

func newRandImage(x, y int) *Image {
	m := Image{Width:100, Height: 100}
	m.colours = make([][]color.RGBA, m.Height)
	for x := range m.colours {
		m.colours[x] = make([]color.RGBA, m.Width)
		for y := range m.colours[x] {
			m.colours[x][y] = color.RGBA{uint8(rand.Intn(255)),
				uint8(rand.Intn(255)),
				uint8(rand.Intn(255)),
				uint8(rand.Intn(255))}
		}
	}
	return &m
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// m is a pointer
	m := newRandImage(100, 100)
	pic.ShowImage(m)
}
