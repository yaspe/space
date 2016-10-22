package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"math/rand"
)

type Vertex struct {
	x float64
	y float64
}

type AbstractObject struct {
	mass           uint64
	size           uint
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

	inFramePosition := &Vertex{
		x: frameSize / 2,
		y: frameSize / 2,
	}

	renderer.SetDrawColor(225, 20, 20, 255)
	const flame_len = 15
	if s.acceleration.x != 0 || s.acceleration.y != 0 {
		for i := 0; i <= 50; i++ {
			renderer.DrawLine(
				int(inFramePosition.x),
				int(inFramePosition.y),
				int(inFramePosition.x-flame_len*math.Cos(s.rotation))+5-rand.Intn(10),
				int(inFramePosition.y-flame_len*math.Sin(s.rotation))+5-rand.Intn(10),
			)
		}
	}

	if s.rotation_acc != 0 {
		for i := 0; i <= 50; i++ {
			if s.rotation_acc < 0 {
				renderer.DrawLine(
					int(inFramePosition.x+half_size*math.Cos(s.rotation+3*math.Pi/4)),
					int(inFramePosition.y+half_size*math.Sin(s.rotation+3*math.Pi/4)),
					int(inFramePosition.x+half_size*math.Cos(s.rotation+3*math.Pi/4))+5-rand.Intn(10),
					int(inFramePosition.y+half_size*math.Sin(s.rotation+3*math.Pi/4))+5-rand.Intn(10),
				)
			} else {
				renderer.DrawLine(
					int(inFramePosition.x-half_size*math.Sin(s.rotation+3*math.Pi/4)),
					int(inFramePosition.y+half_size*math.Cos(s.rotation+3*math.Pi/4)),
					int(inFramePosition.x-half_size*math.Sin(s.rotation+3*math.Pi/4))+5-rand.Intn(10),
					int(inFramePosition.y+half_size*math.Cos(s.rotation+3*math.Pi/4))+5-rand.Intn(10),
				)
			}
		}
	}

	renderer.SetDrawColor(9, 22, 79, 255)
	for i := 0; i <= 50; i++ {
		renderer.DrawLine(
			int(inFramePosition.x+half_size*math.Cos(s.rotation)),
			int(inFramePosition.y+half_size*math.Sin(s.rotation)),
			int(inFramePosition.x-float64(i)*half_size*math.Sin(s.rotation+3*math.Pi/4)/50),
			int(inFramePosition.y+float64(i)*half_size*math.Cos(s.rotation+3*math.Pi/4)/50),
		)
		renderer.DrawLine(
			int(inFramePosition.x+half_size*math.Cos(s.rotation)),
			int(inFramePosition.y+half_size*math.Sin(s.rotation)),
			int(inFramePosition.x+float64(i)*half_size*math.Cos(s.rotation+3*math.Pi/4)/50),
			int(inFramePosition.y+float64(i)*half_size*math.Sin(s.rotation+3*math.Pi/4)/50),
		)
	}

	renderer.SetDrawColor(0, 70, 70, 70)
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
	s.rotation_acc = enginePower / 20
}

func (s *Ship) EngineLeftDesable() {
	s.rotation_acc = 0
}

func (s *Ship) EngineRight() {
	s.rotation_acc = -enginePower / 20
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

// class Planet

type Planet struct {
	AbstractObject
}

func (p *Planet) Draw(renderer *sdl.Renderer, s *Ship) {
	inFramePosition := RelalculatePos(&p.position, s)

	if inFramePosition.x+float64(p.size) < 0 || inFramePosition.x-float64(p.size) > frameSize || inFramePosition.y+float64(p.size) < 0 || inFramePosition.y-float64(p.size) > frameSize {
		return
	}

	DrawEllipse(p.size, 0xe3, 0xf3, 0xff, inFramePosition, renderer)
	renderer.SetDrawColor(0, 0, 0, 0)
}
