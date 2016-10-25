package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

const (
	enginePower = 0.02
	maxSpeed    = 2
	worldSize   = 4000
	frameSize   = 720
	G           = 0.005
)

func ProcessControls(s *Ship) {
	var keyLeftPress bool
	var keyRightPress bool
	var keyDownPress bool
	for {
		event := sdl.PollEvent()
		switch t := event.(type) {
		case *sdl.KeyUpEvent:
			if t.Keysym.Sym == 1073741905 {
				keyDownPress = false
			} else if t.Keysym.Sym == 1073741903 {
				keyLeftPress = false
			} else if t.Keysym.Sym == 1073741904 {
				keyRightPress = false
			}
		case *sdl.KeyDownEvent:
			sdl.FlushEvent(sdl.KEYDOWN)
			if t.Keysym.Sym == 1073741905 {
				keyDownPress = true
			} else if t.Keysym.Sym == 1073741903 {
				keyLeftPress = true
			} else if t.Keysym.Sym == 1073741904 {
				keyRightPress = true
			}
		}

		if keyDownPress {
			s.EngineMain()
		} else {
			s.EngineMainDesable()
		}

		if keyLeftPress && !keyRightPress {
			s.EngineLeft()
		} else if !keyLeftPress && keyRightPress {
			s.EngineRight()
		} else {
			s.EngineRightDesable()
		}
	}
}

var objects []IAbstractObject

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	window, _ := sdl.CreateWindow("Omg tittle", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		frameSize, frameSize, sdl.WINDOW_SHOWN)
	defer window.Destroy()
	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer renderer.Destroy()

	sdl.JoystickEventState(sdl.ENABLE)

	objectsBG := []*BgStar{}
	for i := 0; i <= 10000; i++ {
		b := &BgStar{}
		b.position.x = float64(rand.Intn(worldSize))
		b.position.y = float64(rand.Intn(worldSize))
		b.size = uint(rand.Intn(3)) + 1
		objectsBG = append(objectsBG, b)
	}

	for i := 0; i < worldSize; i += 500 {
		p := &Planet{}
		p.size = 50 + uint(rand.Intn(100))
		p.mass = 40000
		p.position.x = float64(rand.Intn(worldSize))
		p.position.y = float64(i)
		p.speed = Vertex{x: float64(100-rand.Intn(200)) / 1000, y: 0}
		p.rotation_speed = float64(5-rand.Intn(10)) / 1000
		objects = append(objects, p)
	}

	s := &Ship{}
	s.mass = 1
	s.position.x = 5
	s.position.y = 5
	objects = append(objects, s)

	go ProcessControls(s)

	for {
		renderer.Clear()
		for i := range objectsBG {
			objectsBG[i].Draw(renderer, s)
		}
		for i := range objects {
			objects[i].ApplyGravity(s)
			objects[i].Process()
			objects[i].Draw(renderer, s)
		}

		renderer.Present()
		sdl.Delay(5)
	}
}
