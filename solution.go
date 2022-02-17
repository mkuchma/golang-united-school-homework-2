package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type Side int

const (
	SidesCircle   Side = 0
	SidesTriangle Side = 3
	SidesSquare   Side = 4
)

func CalcSquare(sideLen float64, sidesNum Side) float64 {
	switch sidesNum {
	case Circle:
		return math.Pi * sideLen * sideLen
	case Triangle:
		return sideLen * sideLen * math.Sqrt(3) / 4
	case Square:
		return sideLen * sideLen
	default:
		return 0
	}
}
