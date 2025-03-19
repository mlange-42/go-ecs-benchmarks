package delete2comp

import (
	"testing"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArk(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld(1024)

	mapper := ecs.NewMap2[comps.Position, comps.Velocity](&world)

	entities := make([]ecs.Entity, 0, n)
	mapper.NewBatchFn(n, func(entity ecs.Entity, _ *comps.Position, _ *comps.Velocity) {
		entities = append(entities, entity)
	})

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		for _, e := range entities {
			world.RemoveEntity(e)
		}
		b.StopTimer()

		entities = entities[:0]

		if n < 64 {
			// Speed up entity creation for low entity counts
			for range n {
				e := mapper.NewEntityFn(nil)
				entities = append(entities, e)
			}
		} else {
			mapper.NewBatchFn(n, func(entity ecs.Entity, _ *comps.Position, _ *comps.Velocity) {
				entities = append(entities, entity)
			})
		}
	}
}

func runArkBatched(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld(1024)

	mapper := ecs.NewMap2[comps.Position, comps.Velocity](&world)
	filter := ecs.NewFilter2[comps.Position, comps.Velocity](&world)

	mapper.NewBatchFn(n, nil)

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		world.RemoveEntities(filter.Batch(), nil)
		b.StopTimer()
		mapper.NewBatchFn(n, nil)
	}
}
