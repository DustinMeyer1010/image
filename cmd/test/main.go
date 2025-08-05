package main

import (
	"image"
	"image/color"

	"github.com/DustinMeyer1010/converters/internal/image/png"
	"github.com/hajimehoshi/ebiten/v2"
)

var screenImage *ebiten.Image

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(screenImage, nil)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 1600, 1600 // Set the canvas size
}

func main() {
	width, height := 234, 215
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))

	image, _ := png.CreatePNG("/Users/dustinmeyer/Documents/Github/image/internal/image/testimages/peter.png")

	data := image.DecodePNG()

	// Set each pixel manually (e.g. checkerboard pattern)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			//rgba.Set(x, y, color.Black)
			offset := y*800 + x
			rgba.Set(x, y, color.RGBA{uint8(data[offset][0]), uint8(data[offset][0]), uint8(data[offset][0]), uint8(data[offset][0])})
			//rgba.Set(x, y, color.RGBA{uint8(data[x][0]), uint8(data[x][1]), uint8(data[x][2]), uint8(data[x][3])}) // red

		}
	}

	screenImage = ebiten.NewImageFromImage(rgba)

	ebiten.SetWindowSize(width*2, height*2)
	ebiten.SetWindowTitle("Draw Pixels")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
