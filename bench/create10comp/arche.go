package create10comp

import (
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArche(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld(1024)

	ids := []ecs.ID{
		ecs.ComponentID[comps.C1](&world),
		ecs.ComponentID[comps.C2](&world),
		ecs.ComponentID[comps.C3](&world),
		ecs.ComponentID[comps.C4](&world),
		ecs.ComponentID[comps.C5](&world),
		ecs.ComponentID[comps.C6](&world),
		ecs.ComponentID[comps.C7](&world),
		ecs.ComponentID[comps.C8](&world),
		ecs.ComponentID[comps.C9](&world),
		ecs.ComponentID[comps.C10](&world),
	}

	ecs.NewBuilder(&world, ids...).NewBatch(n)
	world.Batch().RemoveEntities(ecs.All(ids...))

	entities := make([]ecs.Entity, 0, n)

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		for range n {
			e := world.NewEntity(ids...)
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
			world.Batch().RemoveEntities(ecs.All(ids...))
		}

		entities = entities[:0]
	}
}

func runArcheBatched(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld(1024)

	ids := []ecs.ID{
		ecs.ComponentID[comps.C1](&world),
		ecs.ComponentID[comps.C2](&world),
		ecs.ComponentID[comps.C3](&world),
		ecs.ComponentID[comps.C4](&world),
		ecs.ComponentID[comps.C5](&world),
		ecs.ComponentID[comps.C6](&world),
		ecs.ComponentID[comps.C7](&world),
		ecs.ComponentID[comps.C8](&world),
		ecs.ComponentID[comps.C9](&world),
		ecs.ComponentID[comps.C10](&world),
	}

	ecs.NewBuilder(&world, ids...).NewBatch(n)
	world.Batch().RemoveEntities(ecs.All(ids...))

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		ecs.NewBuilder(&world, ids...).NewBatch(n)
		b.StopTimer()
		world.Batch().RemoveEntities(ecs.All(ids...))
	}
}
