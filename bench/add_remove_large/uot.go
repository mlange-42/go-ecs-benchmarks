package addremovelarge

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/unitoftime/ecs"
)

func runUot(b *testing.B, n int) {
	world := ecs.NewWorld()

	queryPos := ecs.Query1[comps.Position](world)
	queryPosVel := ecs.Query2[comps.Position, comps.Velocity](world)
	comp := ecs.C(comps.Velocity{})

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
	}

	entities := make([]ecs.Id, 0, n)

	// Iterate once for more fairness
	queryPos.MapId(func(id ecs.Id, pos *comps.Position) {
		entities = append(entities, id)
	})

	for _, e := range entities {
		ecs.Write(world, e,
			ecs.C(comps.Velocity{}),
		)
	}

	entities = entities[:0]

	queryPosVel.MapId(func(id ecs.Id, pos *comps.Position, vel *comps.Velocity) {
		entities = append(entities, id)
	})

	for _, e := range entities {
		ecs.DeleteComponent(world, e, comp)
	}

	entities = entities[:0]

	for b.Loop() {
		queryPos.MapId(func(id ecs.Id, pos *comps.Position) {
			entities = append(entities, id)
		})

		for _, e := range entities {
			ecs.Write(world, e,
				ecs.C(comps.Velocity{}),
			)
		}

		entities = entities[:0]

		queryPosVel.MapId(func(id ecs.Id, pos *comps.Position, vel *comps.Velocity) {
			entities = append(entities, id)
		})

		for _, e := range entities {
			ecs.DeleteComponent(world, e, comp)
		}

		entities = entities[:0]
	}
}
