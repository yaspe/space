package core

import (
	"math/rand"
	"time"
)

type Space struct {
	ships []IShip
	planets []*Planet
	ticker *time.Ticker
}

func CreteSpace() *Space {
	s := new(Space)
	s.initPlanets()
	s.ticker = time.NewTicker(time.Millisecond)
	go s.runProcess()
	return s
}

func (s *Space) GetPlanets() *[]*Planet {
	return &s.planets
}

func (s *Space) GetShips() *[]IShip {
	return &s.ships
}

func (s *Space) AddShip(ship IShip) {
	s.ships = append(s.ships, ship)
}

func (s *Space) initPlanets() {
	for i := 0; i < worldSize; i += 500 {
		p := &Planet{}
		p.Size = 50 + uint(rand.Intn(100))
		p.mass = 40000
		p.position = &Vertex{X: float64(rand.Intn(worldSize)), Y: float64(i)}
		p.speed = Vertex{X: float64(100-rand.Intn(200)) / 1000, Y: 0}
		p.rotation_speed = float64(5-rand.Intn(10)) / 1000
		s.planets = append(s.planets, p)
	}
}

func (s *Space) runProcess() {
	for range s.ticker.C {
		s.process()
	}
}

func (s *Space) process() {
	for _, obj := range s.ships {
		obj.ApplyGravity(&s.planets)
		obj.Process()
	}
}