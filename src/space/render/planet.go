package render

import (
	"math"
	
	"../core"
)

func (dr *displayRender) DrawPlanets(p *core.Planet) {
	inFramePosition := core.RelalculatePos(p.GetPosition(), dr.AlignPosition)
	
	if inFramePosition.X + float64(p.Size) < 0 || inFramePosition.X - float64(p.Size) > dr.frameSizeX || inFramePosition.Y + float64(p.Size) < 0 || inFramePosition.Y - float64(p.Size) > dr.frameSizeY {
		return
	}
	
	c := color{0xe3, 0xf3, 0xff}
	
	dr.DrawEllipse(p.Size, c, inFramePosition)
	
	craterPosition := inFramePosition
	craterPosition.X = inFramePosition.X + float64(p.Size)/2*math.Sin(p.GetRotation())
	craterPosition.Y = inFramePosition.Y + float64(p.Size)/2*math.Cos(p.GetRotation())
	
	bc := color{0xe3, 0xf3, 0xff}
	
	dr.DrawEllipse(3, bc, craterPosition)
	
	dr.renderer.SetDrawColor(0, 0, 0, 0)
}