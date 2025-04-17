package global

import (
	. "interstonar/internal/utils"
	"math"
)

func CalculateGravitationalForce(a, b Body) Vector3 {
	dist := Sub(b.Position, a.Position)
	distance := math.Sqrt(dist.X*dist.X + dist.Y*dist.Y + dist.Z*dist.Z)
	if distance < 0 {
		return Vector3{}
	}

	forceMagnitude := G * a.Mass * b.Mass / (distance * distance)
	direction := Vector3{
		X: dist.X / distance,
		Y: dist.Y / distance,
		Z: dist.Z / distance,
	}
	force := Vector3{
		X: direction.X * forceMagnitude,
		Y: direction.Y * forceMagnitude,
		Z: direction.Z * forceMagnitude,
	}
	return force
}

func UpdateBody(body Body, force Vector3, dt float64) Body {
	acceleration := Vector3{
		X: force.X / body.Mass,
		Y: force.Y / body.Mass,
		Z: force.Z / body.Mass,
	}

	body.Velocity.X += acceleration.X * dt
	body.Velocity.Y += acceleration.Y * dt
	body.Velocity.Z += acceleration.Z * dt

	body.Position.X += body.Velocity.X * dt
	body.Position.Y += body.Velocity.Y * dt
	body.Position.Z += body.Velocity.Z * dt

	return body
}
