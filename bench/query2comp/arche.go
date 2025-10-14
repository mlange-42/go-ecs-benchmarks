package query2comp

import (
	"runtime"
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

	loop := func() {
		query := world.Query(&filter)
		for query.Next() {
			pos := (*comps.Position)(query.Get(posID))
			vel := (*comps.Velocity)(query.Get(velID))
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}

	for b.Loop() {
		loop()
	}

	sum := 0.0
	query = world.Query(&filter)
	for query.Next() {
		pos := (*comps.Position)(query.Get(posID))
		sum += pos.X + pos.Y
	}
	runtime.KeepAlive(sum)
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

	loop := func() {
		query := world.Query(&cf)
		for query.Next() {
			pos := (*comps.Position)(query.Get(posID))
			vel := (*comps.Velocity)(query.Get(velID))
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}

	for b.Loop() {
		loop()
	}

	sum := 0.0
	query = world.Query(&cf)
	for query.Next() {
		pos := (*comps.Position)(query.Get(posID))
		sum += pos.X + pos.Y
	}
	runtime.KeepAlive(sum)
}
