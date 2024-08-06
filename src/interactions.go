package main

import(
    "github.com/hajimehoshi/ebiten/v2"
	"github.com/Valdotorium/gobird/pkg/touch"
)
type Touch struct{
	position Vector2i
	release bool
}

type Mouse struct {
	isDown bool
	position Vector2i
}

func GetTouches()*Touch{
	touches := touch.GetTouchIDs()
	for i := range touches{
		touchposx, touchposy := ebiten.TouchPosition(touches[i])
		return &Touch{
			position :  Vector2i{x:touchposx, y:touchposy},
            release : touch.IsTouchJustReleased(touches[i])}
	}
	return nil
}

func UpdateMouse(g *Game)*Game{
	touched := GetTouches()
    if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) || touched != nil {
        g.Mouse.isDown = true
    }
	mouseposx, mouseposy := ebiten.CursorPosition()
	if touched != nil{
		g.Mouse.position = touched.position
	} else {
		g.Mouse.position = Vector2i{x:mouseposx, y:mouseposy}
	}

    return g
}