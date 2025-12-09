package int_point_3d

import "math"

type Location struct {
	X, Y, Z int
}

func (l Location) DistanceTo(other Location) float64 {
	xDelta := l.X - other.X
	yDelta := l.Y - other.Y
	zDelta := l.Z - other.Z
	return math.Sqrt(float64(xDelta*xDelta + yDelta*yDelta + zDelta*zDelta))
}

func At(x, y, z int) Location {
	return Location{X: x, Y: y, Z: z}
}

func Compare(l *Location, r *Location) int {
	if l.X != r.X {
		return l.X - r.X
	} else if l.Y != r.Y {
		return l.Y - r.Y
	}
	return l.Z - r.Z
}
