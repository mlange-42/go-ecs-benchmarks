package create2compalloc

import (
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArche(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		world := ecs.NewWorld(ecs.NewConfig().WithCapacityIncrement(1024))

		posID := ecs.ComponentID[comps.Position](&world)
		velID := ecs.ComponentID[comps.Velocity](&world)
		ids := []ecs.ID{posID, velID}

		b.StartTimer()
		for range n {
			world.NewEntity(ids...)
		}
		b.StopTimer()
	}
}

func runArcheBatched(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		world := ecs.NewWorld(ecs.NewConfig().WithCapacityIncrement(1024))

		posID := ecs.ComponentID[comps.Position](&world)
		velID := ecs.ComponentID[comps.Velocity](&world)
		ids := []ecs.ID{posID, velID}

		b.StartTimer()
		ecs.NewBuilder(&world, ids...).NewBatch(n)
		b.StopTimer()
	}
}
