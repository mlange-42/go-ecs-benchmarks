package addremove

import (
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArche(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld(1024)

	posID := ecs.ComponentID[comps.Position](&world)
	velID := ecs.ComponentID[comps.Velocity](&world)
	ids := []ecs.ID{velID}

	ecs.NewBuilder(&world, posID).NewBatch(n)

	filterPos := ecs.All(posID)
	filterPosVel := ecs.All(posID, velID)

	entities := make([]ecs.Entity, 0, n)

	// Iterate once for more fairness
	query := world.Query(&filterPos)
	for query.Next() {
		entities = append(entities, query.Entity())
	}

	for _, e := range entities {
		world.Add(e, ids...)
	}

	entities = entities[:0]
	query = world.Query(&filterPosVel)
	for query.Next() {
		entities = append(entities, query.Entity())
	}

	for _, e := range entities {
		world.Remove(e, ids...)
	}

	entities = entities[:0]

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		query := world.Query(&filterPos)
		for query.Next() {
			entities = append(entities, query.Entity())
		}

		for _, e := range entities {
			world.Add(e, ids...)
		}

		entities = entities[:0]
		query = world.Query(&filterPosVel)
		for query.Next() {
			entities = append(entities, query.Entity())
		}

		for _, e := range entities {
			world.Remove(e, ids...)
		}

		entities = entities[:0]
	}
}

func runArcheBatched(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld(1024)

	posID := ecs.ComponentID[comps.Position](&world)
	velID := ecs.ComponentID[comps.Velocity](&world)
	ids := []ecs.ID{velID}

	ecs.NewBuilder(&world, posID).NewBatch(n)

	filterPos := ecs.All(posID)
	filterPosVel := ecs.All(posID, velID)

	// Iterate once for more fairness
	world.Batch().Add(&filterPos, ids...)
	world.Batch().Remove(&filterPosVel, ids...)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		world.Batch().Add(&filterPos, ids...)
		world.Batch().Remove(&filterPosVel, ids...)
	}
}
