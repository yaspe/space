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

func RelalculatePos(pos *Vertex, s *Ship) *Vertex {
	inFramePosition := &Vertex{
		x: pos.x - s.position.x + frameSize/2,
		y: pos.y - s.position.y + frameSize/2,
	}

	if inFramePosition.x > worldSize {
		inFramePosition.x = inFramePosition.x - worldSize
	}
	if inFramePosition.y > worldSize {
		inFramePosition.y = inFramePosition.y - worldSize
	}
	if s.position.x > worldSize-frameSize/2 && pos.x < frameSize/2 {
		inFramePosition.x = worldSize + pos.x - s.position.x + frameSize/2
	}
	if s.position.y > worldSize-frameSize/2 && pos.y < frameSize/2 {
		inFramePosition.y = worldSize + pos.y - s.position.y + frameSize/2
	}

	return inFramePosition
}
