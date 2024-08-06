package main

import "fmt"

func PlayerPhysics(g *Game)*Game{
	//limiting the players velocity
	if g.PlayerYVelocity > 10 {
		g.PlayerYVelocity = 10
	}else if g.PlayerYVelocity < -10 {
        g.PlayerYVelocity = -10
    }else{
		g.PlayerYVelocity -= 0.3
	}
	g.PlayerY += g.PlayerYVelocity
	//no falling below ground!
	if g.PlayerY < float32(GroundY){
        g.PlayerY = float32(GroundY)
        g.PlayerYVelocity = 0
    }
	//slowly getting faster
	g.PlayerXVelocity += 0.00002
	g.PlayerX += g.PlayerXVelocity
	g.ScrollX = g.PlayerX - 200
	if g.IsDebuggingMode{
        fmt.Println("PlayerY: ", g.PlayerY, "PlayerYVelocity: ", g.PlayerYVelocity)
		fmt.Println("PlayerScrollX: ", g.ScrollX, "PlayerXVelocity: ", g.PlayerXVelocity,"PlayerX: ", g.PlayerX)
	}

	return g

}