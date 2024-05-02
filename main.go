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

//web build instructions:

//currently building or running?
var build bool = true
type Game struct{
	//go dict
	Textures map[string]*ebiten.Image
	PlayerY float32
	PlayerX float32
	ScrollX float32
	PlayerXVelocity float32
	PlayerYVelocity float32
	IsDebuggingMode bool
	Score int
}
func NewGame() *Game {
	return &Game{
        Textures: LoadImages(imagePaths),
		PlayerY: float32(GroundY),
        PlayerX: float32(PlayerX),
		ScrollX: float32(0),
		PlayerXVelocity: float32(4),
        PlayerYVelocity: float32(0),
		IsDebuggingMode: true,
		Score: 0,
    }
}
func (g *Game) Update() error {
	//detecting jumping state
	g = DetectJump(g)
	g = PlayerPhysics(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100,120,180,255})
	ebitenutil.DebugPrint(screen, "Score:")
	//referencing items by image name
	op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(float64(g.PlayerX) - float64(g.ScrollX), float64(WindowHeight)-float64(g.PlayerY))
    screen.DrawImage(g.Textures["assets/stone.png"], op)

	
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