package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

const (
	enginePower = 0.05
	maxSpeed    = 15
	worldSize   = 480
)

func main() {
	window, _ := sdl.CreateWindow("Omg tittle", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		640, 480, sdl.WINDOW_SHOWN)
	defer window.Destroy()
	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer renderer.Destroy()

	s := &Ship{}
	s.position.x = 50
	s.position.y = 50

	for {
		s.Process()
		s.EngineMain()
		renderer.Clear()
		s.Draw(renderer)
		renderer.Present()
		sdl.Delay(50)
	}
}

type Vertex struct {
	x float64
	y float64
}

type AbstractObject struct {
	mass         uint64
	hp           uint
	position     Vertex
	speed        Vertex
	acceleration Vertex
	rotation     float64
}

type IAbstractObject interface {
	Process()
}

type Ship struct {
	AbstractObject
}

func (s *Ship) Draw(renderer *sdl.Renderer) {
	const half_size = 32

	s.rotation = 1.3
	renderer.SetDrawColor(0, 255, 255, 255)
	renderer.DrawLine(
		int(s.position.x),
		int(s.position.y),
		int(s.position.x+half_size*math.Cos(s.rotation)),
		int(s.position.y+half_size*math.Sin(s.rotation)),
	)
	renderer.DrawLine(
		int(s.position.x),
		int(s.position.y),
		int(s.position.x+half_size*math.Cos(s.rotation+3*math.Pi/4)),
		int(s.position.y+half_size*math.Sin(s.rotation+3*math.Pi/4)),
	)
	renderer.DrawLine(
		int(s.position.x),
		int(s.position.y),
		int(s.position.x-half_size*math.Sin(s.rotation+3*math.Pi/4)),
		int(s.position.y+half_size*math.Cos(s.rotation+3*math.Pi/4)),
	)
	renderer.DrawLine(
		int(s.position.x+half_size*math.Cos(s.rotation)),
		int(s.position.y+half_size*math.Sin(s.rotation)),
		int(s.position.x+half_size*math.Cos(s.rotation+3*math.Pi/4)),
		int(s.position.y+half_size*math.Sin(s.rotation+3*math.Pi/4)),
	)
	renderer.DrawLine(
		int(s.position.x+half_size*math.Cos(s.rotation)),
		int(s.position.y+half_size*math.Sin(s.rotation)),
		int(s.position.x-half_size*math.Sin(s.rotation+3*math.Pi/4)),
		int(s.position.y+half_size*math.Cos(s.rotation+3*math.Pi/4)),
	)
	renderer.SetDrawColor(0, 0, 0, 0)
}

func (s *Ship) EngineMain() {
	s.acceleration.x = enginePower * math.Cos(s.rotation)
	s.acceleration.y = enginePower * math.Sin(s.rotation)
}

func (s *AbstractObject) Process() {

	s.speed.x += s.acceleration.x
	s.speed.y += s.acceleration.y

	if s.speed.x > maxSpeed {
		s.speed.x = maxSpeed
	}

	if s.speed.y > maxSpeed {
		s.speed.y = maxSpeed
	}

	s.position.x += s.speed.x
	s.position.y += s.speed.y

	if s.position.x > worldSize {
		s.position.x = 0
	} else if s.position.x < 0 {
		s.position.x = worldSize
	}

	if s.position.y > worldSize {
		s.position.y = 0
	} else if s.position.y < 0 {
		s.position.y = worldSize
	}
}
