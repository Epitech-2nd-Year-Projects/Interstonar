package local

import . "interstonar/internal/utils"

type Shape struct {
	Type        string
	Position    Vector3
	Radius      float64
	Height      float64
	Sides       Vector3
	InnerRadius float64
	OuterRadius float64
}

type RaymarchResult struct {
	Hit      bool
	Position Vector3
	Steps    int
}
