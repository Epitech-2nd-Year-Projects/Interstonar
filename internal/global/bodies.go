package global

import (
	"interstonar/internal/config"
	. "interstonar/internal/utils"
	"math"
	"sort"
	"strings"
)

const (
	G = 6.674e-11
)

func NewBody(config config.GlobalBody) Body {
	return Body{
		Name:     config.Name,
		Position: config.Position,
		Velocity: config.Direction,
		Mass:     config.Mass,
		Radius:   config.Radius,
		IsGoal:   config.Goal,
	}
}

func CheckCollision(a, b Body) bool {
	dist := Distance(a.Position, b.Position)
	return dist <= a.Radius+b.Radius
}

func MergeBodies(bodies []Body, indices []int) Body {
	totalMass := 0.0
	totalVolume := 0.0
	var names []string

	weightedPosition := Vector3{}
	weightedVelocity := Vector3{}
	isGoal := false

	for _, index := range indices {
		body := bodies[index]
		names = append(names, body.Name)
		totalMass += body.Mass
		totalVolume += (4.0 / 3.0) * math.Pi * math.Pow(body.Radius, 3)

		weightedPosition.X += body.Position.X * body.Mass
		weightedPosition.Y += body.Position.Y * body.Mass
		weightedPosition.Z += body.Position.Z * body.Mass

		weightedVelocity.X += body.Velocity.X * body.Mass
		weightedVelocity.Y += body.Velocity.Y * body.Mass
		weightedVelocity.Z += body.Velocity.Z * body.Mass

		if body.IsGoal {
			isGoal = true
		}
	}

	if totalMass > 0 {
		weightedVelocity.X /= totalMass
		weightedVelocity.Y /= totalMass
		weightedVelocity.Z /= totalMass

		weightedPosition.X /= totalMass
		weightedPosition.Y /= totalMass
		weightedPosition.Z /= totalMass
	}

	newRadius := math.Pow((3.0*totalVolume)/(4.0*math.Pi), 1.0/3.0)

	sort.Strings(names)
	newName := strings.Join(names, "")

	return Body{
		Name:     newName,
		Position: weightedPosition,
		Velocity: weightedVelocity,
		Mass:     totalMass,
		Radius:   newRadius,
		IsGoal:   isGoal,
	}
}
