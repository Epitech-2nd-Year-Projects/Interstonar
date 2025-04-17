package local

import (
	"interstonar/internal/config"
)

const (
	ShapeSphere   = "sphere"
	ShapeCylinder = "cylinder"
	ShapeBox      = "box"
	ShapeTorus    = "torus"
)

func NewShape(config config.LocalShape) Shape {
	return Shape{
		Type:        config.Type,
		Position:    config.Position,
		Radius:      config.Radius,
		Height:      config.Height,
		Sides:       config.Sides,
		InnerRadius: config.InnerRadius,
		OuterRadius: config.OuterRadius,
	}
}
