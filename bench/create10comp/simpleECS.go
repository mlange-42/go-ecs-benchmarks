package create10comp

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	// initialize arrays to N length to avoid memory allocs
	world := ecs.New(n)
	ecs.Register2[comps.Position, comps.Velocity](world)
	ecs.Register5[comps.C1, comps.C2, comps.C3, comps.C4, comps.C5](world)
	ecs.Register5[comps.C6, comps.C7, comps.C8, comps.C9, comps.C10](world)

	entities := make([]ecs.Entity, 0, n)
	for b.Loop() {
		for range n {
			e := ecs.NewEntity(world)
			ecs.Add3(world, e,
				comps.Position{},
				comps.Velocity{X: 1, Y: 1},
				comps.C10{},
			)
			ecs.Add3(world, e, comps.C1{}, comps.C2{}, comps.C3{})
			ecs.Add3(world, e, comps.C4{}, comps.C5{}, comps.C6{})
			ecs.Add3(world, e, comps.C7{}, comps.C8{}, comps.C9{})
			entities = append(entities, e)
		}
	}
	ecs.Kill(world, entities...)
	entities = entities[:0]
}
