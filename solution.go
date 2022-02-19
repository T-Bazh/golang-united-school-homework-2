package square

import "math"

type SidesNumInt int

const (
	SidesCircle   SidesNumInt = 0
	SidesTriangle SidesNumInt = 3
	SidesSquare   SidesNumInt = 4
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum SidesNumInt) float64 {
	switch sidesNum {
	case SidesSquare:
		return sideLen * sideLen
	case SidesTriangle:
		return sideLen * sideLen * math.Sqrt(3) / 4
	case SidesCircle:
		return sideLen * sideLen * math.Pi
	}
	return 0
}
