package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


var build bool = true
type Game struct{
	images []*ebiten.Image
}
func NewGame() *Game {
	return &Game{
        images: LoadImages(),
    }
}
func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	xpos := 0
	for _, image := range g.images {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(xpos), 0)
		screen.DrawImage(image, op)
        xpos += image.Bounds().Size().X
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}