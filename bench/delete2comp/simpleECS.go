package delete2comp

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.New(1024).EnableGrowing()
	ecs.Register2[comps.Position, comps.Velocity](world)
	for range n {
		e := ecs.NewEntity(world)
		ecs.Add2(world, e,
			comps.Position{},
			comps.Velocity{X: 1, Y: 1},
		)
	}
	stPosition, stVelocity := ecs.GetStorage2[comps.Position, comps.Velocity](world)
	for b.Loop() {
		entities := stPosition.And(stVelocity)
		b.StartTimer()
		for _, e := range entities {
			ecs.Kill(world, e)
		}
		b.StopTimer()
		for range n {
			e := ecs.NewEntity(world)
			ecs.Add2(world, e,
				comps.Position{},
				comps.Velocity{X: 1, Y: 1},
			)
		}
	}
}
// modify pointers instead of copying components
func runSimpleECS_Batch(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.New(1024).EnableGrowing()
	ecs.Register2[comps.Position, comps.Velocity](world)
	for range n {
		e := ecs.NewEntity(world)
		ecs.Add2(world, e,
			comps.Position{},
			comps.Velocity{X: 1, Y: 1},
		)
	}
	for b.Loop() {
		stPosition, stVelocity := ecs.GetStorage2[comps.Position, comps.Velocity](world)
		e := stPosition.And(stVelocity)
		b.StartTimer()
		ecs.Kill(world, e...)
		b.StopTimer()

		for range n {
			e := ecs.NewEntity(world)
			ecs.Add2(world, e,
				comps.Position{},
				comps.Velocity{X: 1, Y: 1},
			)
		}
	}
}
