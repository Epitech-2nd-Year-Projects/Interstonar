package global

import (
	"fmt"
	"interstonar/internal/config"
	. "interstonar/internal/utils"
)

const (
	DeltaTime = 3600.0
	MaxDays   = 365
	MaxSteps  = MaxDays * 24
)

func Simulate(conf *config.GlobalConfig, position, velocity []float64) {
	rockPos := Vector3{
		X: position[0],
		Y: position[1],
		Z: position[2],
	}
	rockVel := Vector3{
		X: velocity[0],
		Y: velocity[1],
		Z: velocity[2],
	}

	rock := Rock{
		Position: rockPos,
		Velocity: rockVel,
		Mass:     1.0,
	}

	var bodies []Body
	for _, bodyConf := range conf.Bodies {
		bodies = append(bodies, NewBody(bodyConf))
	}

	for step := 1; step <= MaxSteps; step++ {
		fmt.Printf("At time t = %d: rock is (%.3f, %.3f, %.3f)\n", step, rock.Position.X, rock.Position.Y, rock.Position.Z)

		for _, body := range bodies {
			rockBody := Body{
				Position: rock.Position,
				Radius:   0.1,
			}

			if CheckCollision(rockBody, body) {
				fmt.Printf("Collision between rock and %s\n", body.Name)
				if body.IsGoal {
					fmt.Println("Mission success")
				} else {
					fmt.Println("Mission failure")
				}
				return
			}
		}

		rockForce := Vector3{}
		for _, body := range bodies {
			rockBody := Body{
				Position: rock.Position,
				Mass:     rock.Mass,
			}

			force := CalculateGravitationalForce(rockBody, body)
			rockForce.X += force.X
			rockForce.Y += force.Y
			rockForce.Z += force.Z
		}

		rockAcceleration := Vector3{
			X: rockForce.X / rock.Mass,
			Y: rockForce.Y / rock.Mass,
			Z: rockForce.Z / rock.Mass,
		}

		rock.Velocity.X += rockAcceleration.X * DeltaTime
		rock.Velocity.Y += rockAcceleration.Y * DeltaTime
		rock.Velocity.Z += rockAcceleration.Z * DeltaTime

		rock.Position.X += rock.Velocity.X * DeltaTime
		rock.Position.Y += rock.Velocity.Y * DeltaTime
		rock.Position.Z += rock.Velocity.Z * DeltaTime

		var newBodies []Body
		var bodiesInCollision = make(map[int]bool)

		for i := 0; i < len(bodies); i++ {
			if bodiesInCollision[i] {
				continue
			}

			var collidingBodies []int
			collidingBodies = append(collidingBodies, i)

			for j := i + 1; j < len(bodies); j++ {
				if bodiesInCollision[j] {
					continue
				}

				if CheckCollision(bodies[i], bodies[j]) {
					collidingBodies = append(collidingBodies, j)
					bodiesInCollision[j] = true
				}
			}

			if len(collidingBodies) > 1 {
				bodiesInCollision[i] = true
				newBody := MergeBodies(bodies, collidingBodies)
				newBodies = append(newBodies, newBody)
			} else if !bodiesInCollision[i] {
				totalForce := Vector3{}
				for j, otherBody := range bodies {
					if i != j && !bodiesInCollision[j] {
						force := CalculateGravitationalForce(bodies[i], otherBody)
						totalForce.X += force.X
						totalForce.Y += force.Y
						totalForce.Z += force.Z
					}
				}
				updatedBody := bodies[i]
				updatedBody = UpdateBody(updatedBody, totalForce, DeltaTime)
				newBodies = append(newBodies, updatedBody)
			}
		}
		bodies = newBodies
	}
	fmt.Println("Mission failure")
}
