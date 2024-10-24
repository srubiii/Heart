package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Heart struct {
	screenWidth  int
	screenHeight int
}

func (h *Heart) Update() error {
	return nil
}

func (h *Heart) DrawHalfCircle(screen *ebiten.Image, cx, cy, radius float64, color color.RGBA) {
	// Draw a half-circle using small lines
	segments := 30.0        // The number of segments to approximate the curve (converted to float64)
	startAngle := math.Pi   // Start at 180 degrees (left)
	endAngle := 2 * math.Pi // End at 360degrees (right)

	for i := 0; i < int(segments); i++ {
		theta1 := startAngle + (endAngle-startAngle)*float64(i)/segments
		theta2 := startAngle + (endAngle-startAngle)*float64(i+1)/segments

		x1 := cx + radius*math.Cos(theta1)
		y1 := cy + radius*math.Sin(theta1)
		x2 := cx + radius*math.Cos(theta2)
		y2 := cy + radius*math.Sin(theta2)

		vector.StrokeLine(screen, float32(x1), float32(y1), float32(x2), float32(y2), 2, color, false)
	}
}

func (h *Heart) Draw(screen *ebiten.Image) {
	// Draw heart shape
	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}

	// Draw half-circles for the top part of the heart
	h.DrawHalfCircle(screen, 100, 100, 50, red) // Left half-circle
	h.DrawHalfCircle(screen, 200, 100, 50, red) // Right half-circle

	// Draw the lines to form the bottom part of the heart
	vector.StrokeLine(screen, 50, 100, 150, 210, 2, red, false)  // Left line
	vector.StrokeLine(screen, 250, 100, 150, 210, 2, red, false) // Right line

}

func (h *Heart) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 300, 300
}

func main() {
	log.Printf("main: Hello World")
	fmt.Println("Hello There")

	ebiten.SetWindowSize(300, 300)
	ebiten.SetWindowTitle("Heart Shape")
	if err := ebiten.RunGame(&Heart{}); err != nil {
		log.Fatal(err)
	}
}
