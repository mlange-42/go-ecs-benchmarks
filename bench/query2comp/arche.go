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

	for b.Loop() {
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

	for b.Loop() {
		query := world.Query(&cf)
		for query.Next() {
			pos := (*comps.Position)(query.Get(posID))
			vel := (*comps.Velocity)(query.Get(velID))
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
}
