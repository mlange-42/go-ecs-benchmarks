package delete10comp

import (
	"testing"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArk(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	mapper := ecs.NewMap10[
		comps.C1, comps.C2, comps.C3, comps.C4, comps.C5,
		comps.C6, comps.C7, comps.C8, comps.C9, comps.C10,
	](world)
	filter := ecs.NewFilter0(world)

	entities := make([]ecs.Entity, 0, n)

	mapper.NewBatchFn(n, nil)
	query := filter.Query()
	for query.Next() {
		entities = append(entities, query.Entity())
	}

	for b.Loop() {
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
			mapper.NewBatchFn(n, nil)
			query := filter.Query()
			for query.Next() {
				entities = append(entities, query.Entity())
			}
		}
		b.StartTimer()
	}
}

func runArkBatched(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	mapper := ecs.NewMap10[
		comps.C1, comps.C2, comps.C3, comps.C4, comps.C5,
		comps.C6, comps.C7, comps.C8, comps.C9, comps.C10,
	](world)
	filter := ecs.NewFilter0(world)

	mapper.NewBatchFn(n, nil)

	for b.Loop() {
		world.RemoveEntities(filter.Batch(), nil)
		b.StopTimer()
		mapper.NewBatchFn(n, nil)
		b.StartTimer()
	}
}
