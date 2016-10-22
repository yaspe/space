package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

func DrawEllipse(radius uint, r, g, b uint8, position *Vertex, renderer *sdl.Renderer) {
	renderer.SetDrawColor(r, g, b, 255)

	for i := 0; i <= 314*2; i++ {
		renderer.DrawLine(
			int(position.x),
			int(position.y),
			int(position.x+float64(radius)*math.Cos(float64(i)/100)),
			int(position.y+float64(radius)*math.Sin(float64(i)/100)),
		)
	}
}
