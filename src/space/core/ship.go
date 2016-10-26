package core

import (
	"math"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_gfx"
)

type Ship struct {
	fisObject
}

func (s *Ship) EngineMain() {
	s.acceleration.X = enginePower * math.Cos(s.rotation)
	s.acceleration.Y = enginePower * math.Sin(s.rotation)
}

func (s *Ship) EngineMainDesable() {
	s.acceleration.X = 0
	s.acceleration.Y = 0
}

func (s *Ship) EngineLeft() {
	s.rotation_acc = enginePower / 100
}

func (s *Ship) EngineLeftDesable() {
	s.rotation_acc = 0
}

func (s *Ship) EngineRight() {
	s.rotation_acc = -enginePower / 100
}

func (s *Ship) EngineRightDesable() {
	s.rotation_acc = 0
}

