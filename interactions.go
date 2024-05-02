package main

import("fmt"
    "github.com/hajimehoshi/ebiten/v2"
)

func DetectJump(){
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) || ebiten.IsMouseButtonPressed(ebiten.MouseButton1) || ebiten.IsKeyPressed(ebiten.KeySpace){
		fmt.Println("jump")
	}
}