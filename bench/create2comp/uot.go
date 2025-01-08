package create2comp

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/unitoftime/ecs"
)

func runUot(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld()

	entities := make([]ecs.Id, 0, n)
	for range n {
		id := world.NewId()
		ecs.Write(world, id,
			ecs.C(comps.Position{}),
			ecs.C(comps.Velocity{}),
		)
		entities = append(entities, id)
	}
	for _, e := range entities {
		ecs.Delete(world, e)
	}
	entities = entities[:0]

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		for range n {
			id := world.NewId()
			ecs.Write(world, id,
				ecs.C(comps.Position{}),
				ecs.C(comps.Velocity{}),
			)
			entities = append(entities, id)
		}
		b.StopTimer()
		for _, e := range entities {
			ecs.Delete(world, e)
		}
		entities = entities[:0]
	}
}
