package local

import (
	"fmt"
	"interstonar/internal/config"
	. "interstonar/internal/utils"
	"math"
)

func Simulate(conf *config.LocalConfig, position, velocity []float64) {
	origin := Vector3{
		X: position[0],
		Y: position[1],
		Z: position[2],
	}
	direction := Vector3{
		X: velocity[0],
		Y: velocity[1],
		Z: velocity[2],
	}

	fmt.Printf("Rock thrown at the point (%.2f, %.2f, %.2f) and parallel to the vector (%.2f, %.2f, %.2f)\n", origin.X, origin.Y, origin.Z, direction.X, direction.Y, direction.Z)

	var shapes []Shape
	for _, shapeConf := range conf.Bodies {
		shape := NewShape(shapeConf)
		shapes = append(shapes, shape)

		switch shape.Type {
		case ShapeSphere:
			fmt.Printf("Sphere of radius %.2f at position (%.2f, %.2f, %.2f)\n", shape.Radius, shape.Position.X, shape.Position.Y, shape.Position.Z)
		case ShapeCylinder:
			if shape.Height > 0 {
				fmt.Printf("Cylinder of radius %.2f and height %.2f at position (%.2f, %.2f, %.2f)\n", shape.Radius, shape.Height, shape.Position.X, shape.Position.Y, shape.Position.Z)
			} else {
				fmt.Printf("Infinite cylinder of radius %.2f at position (%.2f, %.2f, %.2f)\n", shape.Radius, shape.Position.X, shape.Position.Y, shape.Position.Z)
			}
		case ShapeBox:
			fmt.Printf("Box of dimensions (%.2f, %.2f, %.2f) at position (%.2f, %.2f, %.2f)\n", shape.Sides.X, shape.Sides.Y, shape.Sides.Z, shape.Position.X, shape.Position.Y, shape.Position.Z)
		case ShapeTorus:
			fmt.Printf("Torus of inner radius %.2f and outer radius %.2f at position (%.2f, %.2f, %.2f)\n", shape.InnerRadius, shape.OuterRadius, shape.Position.X, shape.Position.Y, shape.Position.Z)
		}
	}

	positions := []Vector3{}
	currentPos := origin
	direction = Normalize(direction)

	for step := 1; step <= MaxSteps; step++ {
		dist := MinDistance(currentPos, shapes)

		currentPos.X += direction.X * dist
		currentPos.Y += direction.Y * dist
		currentPos.Z += direction.Z * dist

		positions = append(positions, currentPos)

		if dist <= IntersectDist {
			break
		}
		if dist >= MaxDistance {
			break
		}
	}
	maxStepsToShow := math.Min(float64(len(positions)), 25)
	for i := 0; i < int(maxStepsToShow); i++ {
		if i == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("Step %d: (%.2f, %.2f, %.2f)\n", i+1, positions[i].X, positions[i].Y, positions[i].Z)
	}

	lastDist := MinDistance(positions[len(positions)-1], shapes)
	if lastDist <= IntersectDist {
		fmt.Println("Result: Intersection")
	} else if len(positions) >= MaxSteps {
		fmt.Println("Result: Time out")
	} else {
		fmt.Println("Result: Out of scene")
	}
}
