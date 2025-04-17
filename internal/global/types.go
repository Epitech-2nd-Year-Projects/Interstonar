package global

import . "interstonar/internal/utils"

type Body struct {
	Name     string
	Position Vector3
	Velocity Vector3
	Mass     float64
	Radius   float64
	IsGoal   bool
}

type Rock struct {
	Position Vector3
	Velocity Vector3
	Mass     float64
}
