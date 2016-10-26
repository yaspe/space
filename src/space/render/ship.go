package render

import (
	"math"
	"math/rand"
	
	"github.com/veandco/go-sdl2/sdl_gfx"
	"../core"
)

func (dr *displayRender) DrawShip(s core.IShip) {
	const half_size = 32
	
	inFramePosition := &core.Vertex{
		X: dr.frameSizeX / 2,
		Y: dr.frameSizeY / 2, // o_0"
	}
	
	dr.renderer.SetDrawColor(225, 20, 20, 255)
	
	const flame_len = 15
		if s.GetAcceleration().X != 0 || s.GetAcceleration().Y != 0 {
		for i := 0; i <= 50; i++ {
			dr.renderer.DrawLine(
				int(inFramePosition.X),
				int(inFramePosition.Y),
				int(inFramePosition.X -flame_len*math.Cos(s.GetRotation()))+5-rand.Intn(10),
				int(inFramePosition.Y -flame_len*math.Sin(s.GetRotation()))+5-rand.Intn(10),
			)
		}
	}
	
	//if s.rotation_acc != 0 {
	//	for i := 0; i <= 50; i++ {
	//		if s.rotation_acc < 0 {
	//			renderer.DrawLine(
	//				int(inFramePosition.X +half_size*math.Cos(s.rotation+3*math.Pi/4)),
	//				int(inFramePosition.Y +half_size*math.Sin(s.rotation+3*math.Pi/4)),
	//				int(inFramePosition.X +half_size*math.Cos(s.rotation+3*math.Pi/4))+5-rand.Intn(10),
	//				int(inFramePosition.Y +half_size*math.Sin(s.rotation+3*math.Pi/4))+5-rand.Intn(10),
	//			)
	//		} else {
	//			renderer.DrawLine(
	//				int(inFramePosition.X -half_size*math.Sin(s.rotation+3*math.Pi/4)),
	//				int(inFramePosition.Y +half_size*math.Cos(s.rotation+3*math.Pi/4)),
	//				int(inFramePosition.X -half_size*math.Sin(s.rotation+3*math.Pi/4))+5-rand.Intn(10),
	//				int(inFramePosition.Y +half_size*math.Cos(s.rotation+3*math.Pi/4))+5-rand.Intn(10),
	//			)
	//		}
	//	}
	//}
	
	gfx.FilledTrigonRGBA(dr.renderer,
		int(inFramePosition.X +half_size*math.Cos(s.GetRotation())),
		int(inFramePosition.Y +half_size*math.Sin(s.GetRotation())),
		int(inFramePosition.X -float64(0)*half_size*math.Sin(s.GetRotation()+3*math.Pi/4)/50),
		int(inFramePosition.Y +float64(0)*half_size*math.Cos(s.GetRotation()+3*math.Pi/4)/50),
		int(inFramePosition.X -float64(50)*half_size*math.Sin(s.GetRotation()+3*math.Pi/4)/50),
		int(inFramePosition.Y +float64(50)*half_size*math.Cos(s.GetRotation()+3*math.Pi/4)/50),
		9, 22, 79, 255,
	)
	
	gfx.FilledTrigonRGBA(dr.renderer,
		int(inFramePosition.X +half_size*math.Cos(s.GetRotation())),
		int(inFramePosition.Y +half_size*math.Sin(s.GetRotation())),
		int(inFramePosition.X +float64(0)*half_size*math.Cos(s.GetRotation()+3*math.Pi/4)/50),
		int(inFramePosition.Y +float64(0)*half_size*math.Sin(s.GetRotation()+3*math.Pi/4)/50),
		int(inFramePosition.X +float64(50)*half_size*math.Cos(s.GetRotation()+3*math.Pi/4)/50),
		int(inFramePosition.Y +float64(50)*half_size*math.Sin(s.GetRotation()+3*math.Pi/4)/50),
		9, 22, 79, 255,
	)
	
	dr.renderer.SetDrawColor(0, 70, 70, 70)
	dr.renderer.DrawLine(
		int(inFramePosition.X),
		int(inFramePosition.Y),
		int(inFramePosition.X +half_size*math.Cos(s.GetRotation())),
		int(inFramePosition.Y +half_size*math.Sin(s.GetRotation())),
	)
	dr.renderer.DrawLine(
		int(inFramePosition.X),
		int(inFramePosition.Y),
		int(inFramePosition.X +half_size*math.Cos(s.GetRotation()+3*math.Pi/4)),
		int(inFramePosition.Y +half_size*math.Sin(s.GetRotation()+3*math.Pi/4)),
	)
	dr.renderer.DrawLine(
		int(inFramePosition.X),
		int(inFramePosition.Y),
		int(inFramePosition.X -half_size*math.Sin(s.GetRotation()+3*math.Pi/4)),
		int(inFramePosition.Y +half_size*math.Cos(s.GetRotation()+3*math.Pi/4)),
	)
	dr.renderer.DrawLine(
		int(inFramePosition.X +half_size*math.Cos(s.GetRotation())),
		int(inFramePosition.Y +half_size*math.Sin(s.GetRotation())),
		int(inFramePosition.X +half_size*math.Cos(s.GetRotation()+3*math.Pi/4)),
		int(inFramePosition.Y +half_size*math.Sin(s.GetRotation()+3*math.Pi/4)),
	)
	dr.renderer.DrawLine(
		int(inFramePosition.X +half_size*math.Cos(s.GetRotation())),
		int(inFramePosition.Y +half_size*math.Sin(s.GetRotation())),
		int(inFramePosition.X -half_size*math.Sin(s.GetRotation()+3*math.Pi/4)),
		int(inFramePosition.Y +half_size*math.Cos(s.GetRotation()+3*math.Pi/4)),
	)
	
	dr.renderer.SetDrawColor(0, 0, 0, 0)
}