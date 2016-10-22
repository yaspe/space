package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

const (
	enginePower = 0.25
	maxSpeed    = 5
	worldSize   = 4000
	frameSize   = 600
)

func ProcessControls(s *Ship) {
	event := sdl.PollEvent()
	switch t := event.(type) {
	case *sdl.KeyUpEvent:
		if t.Keysym.Sym == 1073741905 {
			s.EngineMainDesable()
		} else if t.Keysym.Sym == 1073741903 {
			s.EngineLeftDesable()
		} else if t.Keysym.Sym == 1073741904 {
			s.EngineRightDesable()
		}
	case *sdl.KeyDownEvent:
		sdl.FlushEvent(sdl.KEYDOWN)
		if t.Keysym.Sym == 1073741905 {
			s.EngineMain()
		} else if t.Keysym.Sym == 1073741903 {
			s.EngineLeft()
		} else if t.Keysym.Sym == 1073741904 {
			s.EngineRight()
		}
	}

}

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	window, _ := sdl.CreateWindow("Omg tittle", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		frameSize, frameSize, sdl.WINDOW_SHOWN)
	defer window.Destroy()
	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer renderer.Destroy()

	sdl.JoystickEventState(sdl.ENABLE)

	s := &Ship{}
	s.position.x = 50
	s.position.y = 50

	p1 := &Planet{}
	p1.size = 100
	p1.position.x = 250
	p1.position.y = 250
	p1.speed = Vertex{x: 0.05, y: 0}

	p2 := &Planet{}
	p2.size = 125
	p2.position.x = 1025
	p2.position.y = 700
	p2.speed = Vertex{x: -0.1, y: 0}

	p3 := &Planet{}
	p3.size = 70
	p3.position.x = 3700
	p3.position.y = 2700
	p3.speed = Vertex{x: -0.05, y: 0}

	objects := []IAbstractObject{}

	for i := 0; i <= 10000; i++ {
		b := &BgStar{}
		b.position.x = float64(rand.Intn(worldSize))
		b.position.y = float64(rand.Intn(worldSize))
		b.size = 2
		objects = append(objects, b)
	}

	objects = append(objects, p1)
	objects = append(objects, p2)
	objects = append(objects, p3)
	objects = append(objects, s)

	for {
		ProcessControls(s)
		renderer.Clear()
		for i := range objects {
			objects[i].Process()
			objects[i].Draw(renderer, s)
		}
		renderer.Present()
		sdl.Delay(50)
	}
}
