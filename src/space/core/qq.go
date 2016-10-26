package core

const (
	enginePower = 0.02
	maxSpeed    = 2
	worldSize   = 4000
	frameSize   = 720 // o_0"
	G           = 0.005
)

type Vertex struct {
	X float64
	Y float64
}

type IShip interface {
	SetMainEngineValue(uint8)
	SetEngineLeft(uint8)
	EngineRight(uint8)
}

type IDrawObject interface {
	Draw()
}

func RelalculatePos(p1 ,p2 *Vertex) *Vertex {
	inFramePosition := &Vertex{
		X: p1.X - p2.X + frameSize/2,
		Y: p1.Y - p2.Y + frameSize/2,
	}
	
	if inFramePosition.X > worldSize {
		inFramePosition.X = inFramePosition.X - worldSize
	}
	if inFramePosition.Y > worldSize {
		inFramePosition.Y = inFramePosition.Y - worldSize
	}
	if p2.X > worldSize-frameSize/2 && p1.X < frameSize/2 {
		inFramePosition.X = worldSize + p1.X - p2.X + frameSize/2
	}
	if p2.Y > worldSize-frameSize/2 && p1.Y < frameSize/2 {
		inFramePosition.Y = worldSize + p1.Y - p2.Y + frameSize/2
	}
	
	return inFramePosition
}