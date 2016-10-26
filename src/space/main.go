package main

import (
	"./core"
	"./render"
)

const (
	FrameSizeX = 800
	FrameSizeY = 600
)

func main() {
	
	s := core.CreteSpace()
	
	cursorPosition := &core.Vertex{5, 5}
	
	ship := new(core.Ship)
	ship.SetPosition(cursorPosition)
	
	s.AddShip(ship)
	
	r := render.InitRender(FrameSizeX, FrameSizeY, cursorPosition)
	defer r.Destroy()
	

	go ProcessControls(s)

	r.DrawProcess()
}
