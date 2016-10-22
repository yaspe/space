package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

type BgStar struct {
	AbstractObject
}

func (b *BgStar) Draw(renderer *sdl.Renderer, s *Ship) {
	inFramePosition := &Vertex{
		x: b.position.x - s.position.x + frameSize/2,
		y: b.position.y - s.position.y + frameSize/2,
	}

	if inFramePosition.x > worldSize {
		inFramePosition.x = inFramePosition.x - worldSize
	}
	if inFramePosition.y > worldSize {
		inFramePosition.y = inFramePosition.y - worldSize
	}
	if s.position.x > worldSize-frameSize/2 && b.position.x < frameSize/2 {
		inFramePosition.x = worldSize + b.position.x - s.position.x + frameSize/2
	}
	if s.position.y > worldSize-frameSize/2 && b.position.y < frameSize/2 {
		inFramePosition.y = worldSize + b.position.y - s.position.y + frameSize/2
	}

	if inFramePosition.x < 0 || inFramePosition.x > frameSize || inFramePosition.y < 0 || inFramePosition.y > frameSize {
		return
	}

	k := 120 + rand.Intn(30)
	renderer.SetDrawColor(uint8(k), uint8(k), uint8(k), 255)
	rect := &sdl.Rect{int32(inFramePosition.x), int32(inFramePosition.y), 2, 2}
	renderer.DrawRect(rect)
	renderer.SetDrawColor(0, 0, 0, 0)
}
