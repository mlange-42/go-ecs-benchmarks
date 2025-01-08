package query1in10

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

	ecs.NewBuilder(&world, posID).NewBatch(n * 9)
	ecs.NewBuilder(&world, posID, velID).NewBatch(n)

	filter := ecs.All(posID, velID)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		query := world.Query(&filter)
		for query.Next() {
			pos := (*comps.Position)(query.Get(posID))
			vel := (*comps.Velocity)(query.Get(velID))
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
}

func runArcheRegistered(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld(1024)

	posID := ecs.ComponentID[comps.Position](&world)
	velID := ecs.ComponentID[comps.Velocity](&world)

	ecs.NewBuilder(&world, posID).NewBatch(n * 9)
	ecs.NewBuilder(&world, posID, velID).NewBatch(n)

	filter := ecs.All(posID, velID)
	cf := world.Cache().Register(&filter)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		query := world.Query(&cf)
		for query.Next() {
			pos := (*comps.Position)(query.Get(posID))
			vel := (*comps.Velocity)(query.Get(velID))
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
}
