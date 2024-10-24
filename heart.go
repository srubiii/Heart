package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//Draw heart shape lines manually
	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}

	vector.StrokeLine(screen, 100, 100, 150, 70, 2, red, false)
	vector.StrokeLine(screen, 150, 70, 200, 100, 2, red, false)
	vector.StrokeLine(screen, 200, 100, 150, 150, 2, red, false)
	vector.StrokeLine(screen, 150, 150, 100, 100, 2, red, false)
}

func (g *Game) Layout(outsideWidth, outstideHeight int) (int, int) {
	return 320, 240
}

func main() {
	prefix := "main"
	log.Printf("%s: Hello World\n", prefix)
	fmt.Println("Hello There")

	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Heart Shape with StrokeLine")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

	
}
