package local

import (
	. "interstonar/internal/utils"
	"math"
)

func SdfSphere(p Vector3, center Vector3, radius float64) float64 {
	dx := p.X - center.X
	dy := p.Y - center.Y
	dz := p.Z - center.Z
	return math.Sqrt(dx*dx+dy*dy+dz*dz) - radius
}

func SdfCylinder(p Vector3, center Vector3, radius, height float64) float64 {
	dx := p.X - center.X
	dy := p.Y - center.Y
	dz := math.Abs(p.Z-center.Z) - height/2
	distXY := math.Sqrt(dx*dx+dy*dy) - radius

	if height <= 0 {
		return distXY
	}

	if distXY > 0 && dz > 0 {
		return math.Sqrt(distXY*distXY + dz*dz)
	}
	return math.Max(distXY, dz)
}

func SdfBox(p Vector3, center Vector3, sides Vector3) float64 {
	dx := math.Abs(p.X-center.X) - sides.X/2
	dy := math.Abs(p.Y-center.Y) - sides.Y/2
	dz := math.Abs(p.Z-center.Z) - sides.Z/2

	if dx < 0 && dy < 0 && dz < 0 {
		return math.Max(math.Max(dx, dy), dz)
	}

	dx = math.Max(dx, 0)
	dy = math.Max(dy, 0)
	dz = math.Max(dz, 0)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func SdfTorus(p Vector3, center Vector3, innerRadius, outerRadius float64) float64 {
	px := p.X - center.X
	py := p.Y - center.Y
	pz := p.Z - center.Z

	q := math.Sqrt(px*px+py*py) - innerRadius
	return math.Sqrt(q*q+pz*pz) - outerRadius
}

func MinDistance(p Vector3, shapes []Shape) float64 {
	minDist := math.MaxFloat64

	for _, shape := range shapes {
		var dist float64

		switch shape.Type {
		case ShapeSphere:
			dist = SdfSphere(p, shape.Position, shape.Radius)
		case ShapeCylinder:
			dist = SdfCylinder(p, shape.Position, shape.Radius, shape.Height)
		case ShapeBox:
			dist = SdfBox(p, shape.Position, shape.Sides)
		case ShapeTorus:
			dist = SdfTorus(p, shape.Position, shape.InnerRadius, shape.OuterRadius)
		default:
			continue
		}

		if dist < minDist {
			minDist = dist
		}
	}
	return minDist
}
