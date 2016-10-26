package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_gfx"
)

// types

type Vertex struct {
	x float64
	y float64
}

// funcs

func DrawEllipse(radius uint, r, g, b uint8, position *Vertex, renderer *sdl.Renderer) {
	gfx.FilledCircleRGBA(renderer, int(position.x), int(position.y), int(radius), r, g, b, 0xFF)
}

func Sign(a float64) float64 {
	if a > 0 {
		return 1
	}
	return -1
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
