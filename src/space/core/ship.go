package core

import (
	"math"
)

type Ship struct {
	fisObject
}

func (s *Ship) EngineMain() {
	s.acceleration.X = enginePower * math.Cos(s.rotation)
	s.acceleration.Y = enginePower * math.Sin(s.rotation)
}

func (s *Ship) EngineMainDisable() {
	s.acceleration.X = 0
	s.acceleration.Y = 0
}

func (s *Ship) EngineLeft() {
	s.rotation_acc = enginePower / 100
}

func (s *Ship) EngineLeftDisable() {
	s.rotation_acc = 0
}

func (s *Ship) EngineRight() {
	s.rotation_acc = -enginePower / 100
}

func (s *Ship) EngineRightDisable() {
	s.rotation_acc = 0
}

