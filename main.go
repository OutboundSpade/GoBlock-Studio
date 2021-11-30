package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	textImg *ebiten.Image
)

type Game struct {
	custorX int
	cursorY int
}

//Game update
func (g *Game) Update() error {
	g.custorX, g.cursorY = ebiten.CursorPosition()
	return nil
}

// Draws frame
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	var drawOpts *ebiten.DrawImageOptions = &ebiten.DrawImageOptions{}
	drawOpts.GeoM.Translate(float64(g.custorX), float64(g.cursorY))
	textImg.Fill(color.RGBA{255, 0, 0, 255})

	ebitenutil.DebugPrint(textImg, "Ebiten Demos and stuff!")

	screen.DrawImage(textImg, drawOpts)
}

// sets the pixel layout
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func main() {
	textImg = ebiten.NewImage(200, 30)

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("GoBlock Studio")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
