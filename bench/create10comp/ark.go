package create10comp

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

	mapper.NewBatchFn(n, nil)
	world.RemoveEntities(filter.Batch(), nil)

	entities := make([]ecs.Entity, 0, n)

	for b.Loop() {
		for range n {
			e := mapper.NewEntityFn(nil)
			// Just for fairness, because the others need to do that, too.
			entities = append(entities, e)
		}
		b.StopTimer()

		if n < 64 {
			// Speed up cleanup for low entity counts
			for i := len(entities) - 1; i >= 0; i-- {
				world.RemoveEntity(entities[i])
			}
		} else {
			world.RemoveEntities(filter.Batch(), nil)
		}

		entities = entities[:0]
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
	world.RemoveEntities(filter.Batch(), nil)

	for b.Loop() {
		mapper.NewBatchFn(n, nil)
		b.StopTimer()
		world.RemoveEntities(filter.Batch(), nil)
		b.StartTimer()
	}
}
