package addremove

import (
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArche(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	posID := ecs.ComponentID[comps.Position](&world)
	velID := ecs.ComponentID[comps.Velocity](&world)
	ids := []ecs.ID{velID}

	entities := make([]ecs.Entity, 0, n)
	for range n {
		entities = append(entities, world.NewEntity(posID))
	}
	// Iterate once for more fairness
	for _, e := range entities {
		world.Add(e, ids...)
	}
	for _, e := range entities {
		world.Remove(e, ids...)
	}

	for b.Loop() {
		for _, e := range entities {
			world.Add(e, ids...)
		}
		for _, e := range entities {
			world.Remove(e, ids...)
		}
	}
}

func runArcheBatched(b *testing.B, n int) {
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

	for b.Loop() {
		world.Batch().Add(&filterPos, ids...)
		world.Batch().Remove(&filterPosVel, ids...)
	}
}
