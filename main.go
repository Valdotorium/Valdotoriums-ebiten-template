package main

import (
	"log"
    "fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//fetching the paths of the games images from constants.go
var imagePaths []string = fetchGameImagePaths()


//currently building or running?
var build bool = true
type Game struct{
	//go dict
	Textures map[string]*ebiten.Image
	PlayerY float32
	PlayerX float32
}
func NewGame() *Game {
	return &Game{
        Textures: LoadImages(imagePaths),
		PlayerY: float32(GroundY),
        PlayerX: float32(PlayerX),
    }
}
func (g *Game) Update() error {
	//detecting jumping state
	DetectJump()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	//referencing items by image name
	screen.DrawImage(g.Textures["assets/stone.png"], nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	fmt.Println("fetched image paths: ", imagePaths)
	g := NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}