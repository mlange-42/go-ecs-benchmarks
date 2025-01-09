package delete10comp

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

	entities := make([]ecs.Entity, 0, n)

	query := world.Batch().NewQ(n, ids...)
	for query.Next() {
		entities = append(entities, query.Entity())
	}

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		for _, e := range entities {
			world.RemoveEntity(e)
		}
		b.StopTimer()

		entities = entities[:0]
		query := world.Batch().NewQ(n, ids...)
		for query.Next() {
			entities = append(entities, query.Entity())
		}
	}
}

func runArcheBatched(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld(1024)

	posID := ecs.ComponentID[comps.Position](&world)
	velID := ecs.ComponentID[comps.Velocity](&world)
	ids := []ecs.ID{posID, velID}

	filter := ecs.All(ids...)

	world.Batch().New(n, ids...)

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		world.Batch().RemoveEntities(&filter)
		b.StopTimer()
		world.Batch().New(n, ids...)
	}
}
