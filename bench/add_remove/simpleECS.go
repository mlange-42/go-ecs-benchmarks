package addremove

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	world := ecs.New(n)
	ecs.Register2[comps.Position, comps.Velocity](world)
	for range n {
		e := ecs.NewEntity(world)
		ecs.Add(world, e, comps.Position{})
	}
	stPosition, stVelocity := ecs.GetStorage2[comps.Position, comps.Velocity](world)
	for b.Loop() {
		for _, e := range stPosition.And(nil) {
			ecs.Add(world, e, comps.Velocity{})
		}
		// getting entities again is unnesessary, this is to just make it fair
		for _, e := range stVelocity.And(stPosition) {
			ecs.Remove[comps.Velocity](world, e)
		}
	}
	entities := stPosition.And(stVelocity)
	ecs.Kill(world, entities...)

}
