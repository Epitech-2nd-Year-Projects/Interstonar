package utils

import "math"

type Vector3 struct {
	X, Y, Z float64
}

func Add(a, b Vector3) Vector3 {
	return Vector3{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func Sub(a, b Vector3) Vector3 {
	return Vector3{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func Mul(a, b Vector3) Vector3 {
	return Vector3{
		X: a.X * b.X,
		Y: a.Y * b.Y,
		Z: a.Z * b.Z,
	}
}

func Distance(a, b Vector3) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func Normalize(v Vector3) Vector3 {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	if length == 0 {
		return Vector3{0, 0, 0}
	}
	return Vector3{
		X: v.X / length,
		Y: v.Y / length,
		Z: v.Z / length,
	}
}

func Dot(a, b Vector3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
