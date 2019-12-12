package grid

type Orientation int

const (
	North Orientation = 0
	East  Orientation = 1
	South Orientation = 2
	West  Orientation = 3
)

type RelativeDirection int

const (
	Left  RelativeDirection = 0
	Right RelativeDirection = 1
)

func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

type Point struct {
	X, Y int
}

func Turn(orientation Orientation, direction RelativeDirection) Orientation {
	var e Orientation
	switch direction {
	case Left:
		e = orientation - 1
	case Right:
		e = orientation + 1
	}
	if e < 0 {
		e = e + 4
	}
	return e % 4
}

func Move(p Point, orientation Orientation, distance int) Point {
	switch orientation {
	case North:
		return Point{
			X: p.X,
			Y: p.Y + distance,
		}
	case South:
		return Point{
			X: p.X,
			Y: p.Y - distance,
		}
	case East:
		return Point{
			X: p.X + distance,
			Y: p.Y,
		}
	case West:
		return Point{
			X: p.X - distance,
			Y: p.Y,
		}
	default:
		panic("unknown direction")
	}
}
