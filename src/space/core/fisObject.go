package core

import "math"

type fisObject struct {
	mass           float64
	position       *Vertex
	speed          Vertex
	acceleration   Vertex
	gravity        Vertex
	rotation       float64
	rotation_speed float64
	rotation_acc   float64
}

func (f *fisObject) SetPosition(v *Vertex) {
	f.position = v
}

func (f *fisObject) GetPosition() *Vertex {
	return f.position
}

func (f *fisObject) GetAcceleration() *Vertex {
	return &f.acceleration
}

func (f *fisObject) GetRotation() float64 {
	return f.rotation
}

func (f *fisObject) GetMass() float64 {
	return f.mass
}

func (o *fisObject) ApplyGravity(p *[]*Planet) {

	currPoss := RelalculatePos(o.GetPosition(), o.GetPosition())
	o.gravity.X = 0
	o.gravity.Y = 0
	
	for _, planet := range *p {
		
		objPoss := RelalculatePos(planet.GetPosition(), o.GetPosition())
		
		distanceX := currPoss.X - objPoss.X
		distanceY := currPoss.Y - objPoss.Y
		
		distance := math.Sqrt(math.Pow(distanceX, 2) + math.Pow(distanceY, 2))
		if distance < 100 {
			distance = 100
		}
		
		if distance > frameSize {
			continue
		}
		
		g := G * planet.GetMass() / (distance * distance)
		gx := distanceX * g / distance
		gy := distanceY * g / distance
		o.gravity.X += gx
		o.gravity.Y += gy
	}
}

func (s *fisObject) Process() {

	s.speed.X += s.acceleration.X
	s.speed.Y += s.acceleration.Y
	
	if s.speed.X > maxSpeed {
		s.speed.X = maxSpeed
	}
	
	if s.speed.Y > maxSpeed {
		s.speed.Y = maxSpeed
	}
	
	s.speed.X -= s.gravity.X
	s.speed.Y -= s.gravity.Y
	
	s.position.X += s.speed.X
	s.position.Y += s.speed.Y
	
	if s.position.X > worldSize {
		s.position.X = 0
	} else if s.position.X < 0 {
		s.position.X = worldSize
	}
	
	if s.position.Y > worldSize {
		s.position.Y = 0
	} else if s.position.Y < 0 {
		s.position.Y = worldSize
	}
	
	s.rotation_speed += s.rotation_acc
	s.rotation += s.rotation_speed
}