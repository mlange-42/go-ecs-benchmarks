package addremovelarge

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/unitoftime/ecs"
)

func runUot(b *testing.B, n int) {
	world := ecs.NewWorld()

	comp := ecs.C(comps.Velocity{})

	entities := make([]ecs.Id, 0, n)
	for i := 0; i < n; i++ {
		id := world.NewId()
		ecs.Write(world, id,
			ecs.C(comps.Position{}),
			ecs.C(comps.C1{}),
			ecs.C(comps.C2{}),
			ecs.C(comps.C3{}),
			ecs.C(comps.C4{}),
			ecs.C(comps.C5{}),
			ecs.C(comps.C6{}),
			ecs.C(comps.C7{}),
			ecs.C(comps.C8{}),
			ecs.C(comps.C9{}),
			ecs.C(comps.C10{}),
		)
		entities = append(entities, id)
	}

	// Iterate once for more fairness
	for _, e := range entities {
		ecs.Write(world, e,
			ecs.C(comps.Velocity{}),
		)
	}
	for _, e := range entities {
		ecs.DeleteComponent(world, e, comp)
	}

	for b.Loop() {
		for _, e := range entities {
			ecs.Write(world, e,
				ecs.C(comps.Velocity{}),
			)
		}
		for _, e := range entities {
			ecs.DeleteComponent(world, e, comp)
		}
	}
}
