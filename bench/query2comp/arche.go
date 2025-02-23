package query2comp

import (
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArche(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	posID := ecs.ComponentID[comps.Position](&world)
	velID := ecs.ComponentID[comps.Velocity](&world)

	query := ecs.NewBuilder(&world, posID, velID).NewBatchQ(n)
	for query.Next() {
		vel := (*comps.Velocity)(query.Get(velID))
		vel.X = 1
		vel.Y = 1
	}

	filter := ecs.All(posID, velID)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		query := world.Query(&filter)
		for query.Next() {
			pos := (*comps.Position)(query.Get(posID))
			vel := (*comps.Velocity)(query.Get(velID))
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
	b.StopTimer()

	query = world.Query(&filter)
	for query.Next() {
		pos := (*comps.Position)(query.Get(posID))
		if pos.X == 0 || pos.Y == 0 {
			panic("assertion failed")
		}
	}
}

func runArcheRegistered(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	posID := ecs.ComponentID[comps.Position](&world)
	velID := ecs.ComponentID[comps.Velocity](&world)

	query := ecs.NewBuilder(&world, posID, velID).NewBatchQ(n)
	for query.Next() {
		vel := (*comps.Velocity)(query.Get(velID))
		vel.X = 1
		vel.Y = 1
	}

	filter := ecs.All(posID, velID)
	cf := world.Cache().Register(&filter)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		query = world.Query(&cf)
		for query.Next() {
			pos := (*comps.Position)(query.Get(posID))
			vel := (*comps.Velocity)(query.Get(velID))
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
	b.StopTimer()

	query = world.Query(&cf)
	for query.Next() {
		pos := (*comps.Position)(query.Get(posID))
		if pos.X == 0 {
			panic("assertion failed")
		}
	}
}
