package addremovelarge

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	world := ecs.New(n)
	// arrays will be initialized to n length, eliminating memory allocations
	for range n {
		e := ecs.NewEntity(world)
		ecs.Add3(world, e, comps.Position{}, comps.C1{}, comps.C2{})
		ecs.Add3(world, e, comps.C3{}, comps.C4{}, comps.C5{})
		ecs.Add3(world, e, comps.C6{}, comps.C7{}, comps.C8{})
		ecs.Add2(world, e, comps.C9{}, comps.C10{})
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
