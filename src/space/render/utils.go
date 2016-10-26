package render

import (
	"github.com/veandco/go-sdl2/sdl_gfx"
	
	"../core"
)

type color struct {
	r, g, b uint8
}

func (dr *displayRender) DrawEllipse(radius uint, c color, position *core.Vertex) {
	gfx.FilledCircleRGBA(
		dr.renderer,
		int(position.X), int(position.Y),
		int(radius),
		c.r, c.g, c.b,
		0xFF)
}

func Sign(a float64) float64 {
	if a > 0 {
		return 1
	}
	return -1
}


