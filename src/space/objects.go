package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type Vertex struct {
	x float64
	y float64
}

type AbstractObject struct {
	mass           uint64
	hp             uint
	position       Vertex
	speed          Vertex
	acceleration   Vertex
	rotation       float64
	rotation_speed float64
	rotation_acc   float64
}

type IAbstractObject interface {
	Process()
	Draw(renderer *sdl.Renderer, s *Ship)
}

// class Ship

type Ship struct {
	AbstractObject
}

func (s *Ship) Draw(renderer *sdl.Renderer, ss *Ship) {
	const half_size = 32

	renderer.SetDrawColor(0, 255, 255, 255)

	inFramePosition := &Vertex{
		x: frameSize / 2,
		y: frameSize / 2,
	}

	renderer.DrawLine(
		int(inFramePosition.x),
		int(inFramePosition.y),
		int(inFramePosition.x+half_size*math.Cos(s.rotation)),
		int(inFramePosition.y+half_size*math.Sin(s.rotation)),
	)
	renderer.DrawLine(
		int(inFramePosition.x),
		int(inFramePosition.y),
		int(inFramePosition.x+half_size*math.Cos(s.rotation+3*math.Pi/4)),
		int(inFramePosition.y+half_size*math.Sin(s.rotation+3*math.Pi/4)),
	)
	renderer.DrawLine(
		int(inFramePosition.x),
		int(inFramePosition.y),
		int(inFramePosition.x-half_size*math.Sin(s.rotation+3*math.Pi/4)),
		int(inFramePosition.y+half_size*math.Cos(s.rotation+3*math.Pi/4)),
	)
	renderer.DrawLine(
		int(inFramePosition.x+half_size*math.Cos(s.rotation)),
		int(inFramePosition.y+half_size*math.Sin(s.rotation)),
		int(inFramePosition.x+half_size*math.Cos(s.rotation+3*math.Pi/4)),
		int(inFramePosition.y+half_size*math.Sin(s.rotation+3*math.Pi/4)),
	)
	renderer.DrawLine(
		int(inFramePosition.x+half_size*math.Cos(s.rotation)),
		int(inFramePosition.y+half_size*math.Sin(s.rotation)),
		int(inFramePosition.x-half_size*math.Sin(s.rotation+3*math.Pi/4)),
		int(inFramePosition.y+half_size*math.Cos(s.rotation+3*math.Pi/4)),
	)
	renderer.SetDrawColor(0, 0, 0, 0)
}

func (s *Ship) EngineMain() {
	s.acceleration.x = enginePower * math.Cos(s.rotation)
	s.acceleration.y = enginePower * math.Sin(s.rotation)
}

func (s *Ship) EngineMainDesable() {
	s.acceleration.x = 0
	s.acceleration.y = 0
}

func (s *Ship) EngineLeft() {
	s.rotation_acc = enginePower / 10
}

func (s *Ship) EngineLeftDesable() {
	s.rotation_acc = 0
}

func (s *Ship) EngineRight() {
	s.rotation_acc = -enginePower / 10
}

func (s *Ship) EngineRightDesable() {
	s.rotation_acc = 0
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

	s.rotation_speed += s.rotation_acc
	s.rotation += s.rotation_speed
}

// class Ship

type Planet struct {
	AbstractObject
}

func (p *Planet) Draw(renderer *sdl.Renderer, s *Ship) {
	renderer.SetDrawColor(0xe3, 0xf3, 0xff, 255)
	const r = 100

	inFramePosition := &Vertex{
		x: p.position.x - s.position.x + frameSize/2,
		y: p.position.y - s.position.y + frameSize/2,
	}

	for i := 0; i <= 314*2; i++ {
		renderer.DrawLine(
			int(inFramePosition.x),
			int(inFramePosition.y),
			int(inFramePosition.x+r*math.Cos(float64(i)/100)),
			int(inFramePosition.y+r*math.Sin(float64(i)/100)),
		)
	}
	renderer.SetDrawColor(0, 0, 0, 0)
}
