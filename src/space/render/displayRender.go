package render

import (
	"../core"
	
	"github.com/veandco/go-sdl2/sdl"
	"time"
	"math/rand"
)

const bgStarsCount = 10000

type displayRender struct {
	frameSizeX float64
	frameSizeY float64 // то же жесть какая то
	
	window *sdl.Window
	renderer *sdl.Renderer
	
	space core.Space
	
	AlignPosition *core.Vertex
	
	bgStars []*bgStar
}

func InitRender(frameSizeX, frameSizeY int, ap *core.Vertex) *displayRender {
	
	sdl.Init(sdl.INIT_EVERYTHING)
	
	dr := &displayRender{
		frameSizeX: float64(frameSizeX),
		frameSizeY: float64(frameSizeY),
		AlignPosition: ap,
	}
	
	dr.window, _ = sdl.CreateWindow(
		"Omg mega spase game!!!",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		frameSizeX,
		frameSizeY,
		sdl.WINDOW_SHOWN,
	)
	
	dr.renderer, _ = sdl.CreateRenderer(dr.window, -1, sdl.RENDERER_ACCELERATED)
	
	sdl.JoystickEventState(sdl.ENABLE)
	
	dr.initBGStars()
	
	return dr
}

func (dr *displayRender) initBGStars() {
	//dr.bgStars = make([]*bgStar, bgStarsCount)
	for i := 0; i <= bgStarsCount; i++ {
		b := &bgStar{}
		b.position = &core.Vertex{float64(rand.Intn(int(dr.frameSizeX))), float64(rand.Intn(int(dr.frameSizeX)))} // жесть какая то (Слишком много преобразований)
		b.size = uint(rand.Intn(3)) + 1
		//dr.bgStars[i] = b
		dr.bgStars = append(dr.bgStars, b)
	}
}

func (dr *displayRender) DrawProcess() {
	ticker := time.NewTicker(time.Microsecond * 100)
	for range ticker.C {
		dr.renderer.Clear()
		for _, b := range dr.bgStars {
			dr.drawBgStar(b)
		}
		for _, pl := range *dr.space.GetPlanets() {
			dr.DrawPlanets(pl)
		}
		for _, ship := range *dr.space.GetShips() {
			dr.DrawShip(ship)
		}
		dr.renderer.Present()
	}
}

func (dr *displayRender) Destroy() {
	dr.renderer.Destroy()
	dr.window.Destroy()
}