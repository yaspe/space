package main

import (
	"github.com/veandco/go-sdl2/sdl"
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
	/*if s.position.x > worldSize-frameSize && b.position.x < frameSize {
		inFramePosition.x = frameSize - b.position.x + s.position.x
	}*/

	if inFramePosition.x < 0 || inFramePosition.x > frameSize || inFramePosition.y < 0 || inFramePosition.y > frameSize {
		return
	}

	renderer.SetDrawColor(0xe3, 0xf3, 0xff, 255)
	rect := &sdl.Rect{int32(inFramePosition.x), int32(inFramePosition.y), 2, 2}
	renderer.DrawRect(rect)
	renderer.SetDrawColor(0, 0, 0, 0)
}
