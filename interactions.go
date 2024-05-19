package main

import(
    "github.com/hajimehoshi/ebiten/v2"
)

func DetectJump(g *Game)*Game{
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) || ebiten.IsMouseButtonPressed(ebiten.MouseButton1) || ebiten.IsKeyPressed(ebiten.KeySpace) {
		if g.PlayerY == float32(GroundY){
			g.PlayerYVelocity = 10
		}
	}
	return g
}
type Touch struct{
	x int
	y int
	release bool
}
func GetTouches()*Touch{

}