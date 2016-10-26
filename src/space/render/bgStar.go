package render

import (
	"../core"
	
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

type bgStar struct {
	position *core.Vertex
	size uint
}

func (dr *displayRender) drawBgStar(b *bgStar) {
	inFramePosition := core.RelalculatePos(b.position, dr.AlignPosition)
	
	if inFramePosition.X < 0 || inFramePosition.X > dr.frameSizeX || inFramePosition.Y < 0 || inFramePosition.Y > dr.frameSizeY {
		return
	}
	
	k := 120 + rand.Intn(30)
	dr.renderer.SetDrawColor(uint8(k), uint8(k), uint8(k), 255)
	rect := &sdl.Rect{int32(inFramePosition.X), int32(inFramePosition.Y), int32(b.size), int32(b.size)}
	dr.renderer.DrawRect(rect)
	dr.renderer.DrawLine(int(inFramePosition.X)-1, int(inFramePosition.Y)-1, int(inFramePosition.Y)+1, int(inFramePosition.Y)+1)
	dr.renderer.DrawLine(int(inFramePosition.X)+1, int(inFramePosition.Y)-1, int(inFramePosition.Y)-1, int(inFramePosition.Y)+1)
	dr.renderer.SetDrawColor(0, 0, 0, 0)
}
