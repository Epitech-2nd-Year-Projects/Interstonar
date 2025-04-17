package local

import . "interstonar/internal/utils"

const (
	MaxSteps      = 1000
	MaxDistance   = 1000.0
	IntersectDist = 0.1
)

func Raymarch(origin, direction Vector3, shapes []Shape) RaymarchResult {
	direction = Normalize(direction)
	currentPos := origin

	for step := 1; step <= MaxSteps; step++ {
		dist := MinDistance(currentPos, shapes)
		if dist <= IntersectDist {
			return RaymarchResult{
				Hit:      true,
				Position: currentPos,
				Steps:    step,
			}
		}
		if dist >= MaxDistance {
			return RaymarchResult{
				Hit:      false,
				Position: currentPos,
				Steps:    step,
			}
		}

		currentPos.X += direction.X * dist
		currentPos.Y += direction.Y * dist
		currentPos.Z += direction.Z * dist
	}
	return RaymarchResult{
		Hit:      false,
		Position: currentPos,
		Steps:    MaxSteps,
	}
}
