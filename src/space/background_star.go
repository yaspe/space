package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

type BgStar struct {
	AbstractObject
}

func (b *BgStar) Draw(renderer *sdl.Renderer, s *Ship) {
	inFramePosition := RelalculatePos(&b.position, s)

	if inFramePosition.x < 0 || inFramePosition.x > frameSize || inFramePosition.y < 0 || inFramePosition.y > frameSize {
		return
	}

	k := 120 + rand.Intn(30)
	renderer.SetDrawColor(uint8(k), uint8(k), uint8(k), 255)
	rect := &sdl.Rect{int32(inFramePosition.x), int32(inFramePosition.y), int32(b.size), int32(b.size)}
	renderer.DrawRect(rect)
	renderer.DrawLine(int(inFramePosition.x)-1, int(inFramePosition.y)-1, int(inFramePosition.x)+1, int(inFramePosition.y)+1)
	renderer.DrawLine(int(inFramePosition.x)+1, int(inFramePosition.y)-1, int(inFramePosition.x)-1, int(inFramePosition.y)+1)
	renderer.SetDrawColor(0, 0, 0, 0)
}
