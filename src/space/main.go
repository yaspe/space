package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

const (
	enginePower = 0.005
	maxSpeed    = 15
	worldSize   = 480
)

func main() {
	window, _ := sdl.CreateWindow("Omg tittle", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		640, 480, sdl.WINDOW_SHOWN)
	defer window.Destroy()
	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer renderer.Destroy()
	image, _ := sdl.LoadBMP("/Users/ya-spe/Downloads/space_img.bmp")
	defer image.Free()
	texture, _ := renderer.CreateTextureFromSurface(image)
	defer texture.Destroy()

	src := sdl.Rect{0, 0, 64, 64}

	s := &Ship{}

	for {
		s.Process()
		s.engineMain()
		renderer.Clear()
		dst := sdl.Rect{int32(s.position.x), int32(s.position.y), 64, 64}
		renderer.Copy(texture, &src, &dst)
		renderer.Present()
		sdl.Delay(100)
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

func (s *Ship) engineMain() {
	s.acceleration.x = enginePower * math.Sin(s.rotation)
	s.acceleration.y = enginePower * math.Cos(s.rotation)
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

	//if s.position.x
}
