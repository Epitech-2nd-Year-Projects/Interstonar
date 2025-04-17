package config

import (
	. "interstonar/internal/utils"
)

type GlobalBody struct {
	Name      string
	Position  Vector3
	Direction Vector3
	Mass      float64
	Radius    float64
	Goal      bool
}

type GlobalConfig struct {
	Bodies []GlobalBody
}

type LocalShape struct {
	Type        string
	Position    Vector3
	Radius      float64
	Height      float64
	Sides       Vector3
	InnerRadius float64
	OuterRadius float64
}

type LocalConfig struct {
	Bodies []LocalShape
}
