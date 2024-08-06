package main

import (
	"fmt"
	"image/color"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//fetching the paths of the games images from constants.go
var imagePaths []string = fetchGameImagePaths()

type Game struct{
	//go dict
	Textures map[string]*ebiten.Image
	IsDebuggingMode bool
	Score int
	Mouse Mouse
}
func NewGame() *Game {
	return &Game{
        Textures: LoadImages(imagePaths),
		IsDebuggingMode: true,
		Score: 0,
		Mouse: Mouse{},
    }
}
func (g *Game) Update() error {
	//detecting jumping state
	g = UpdateMouse(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100,120,180,255})
	ebitenutil.DebugPrint(screen, "Ebitengine Program")
	
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WindowWidth, WindowHeight
}

func main() {
	fmt.Println("fetched image paths: ", imagePaths)
	g := NewGame()
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}