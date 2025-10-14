package query256arch

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
	c1ID := ecs.ComponentID[comps.C1](&world)
	c2ID := ecs.ComponentID[comps.C2](&world)
	c3ID := ecs.ComponentID[comps.C3](&world)
	c4ID := ecs.ComponentID[comps.C4](&world)
	c5ID := ecs.ComponentID[comps.C5](&world)
	c6ID := ecs.ComponentID[comps.C6](&world)
	c7ID := ecs.ComponentID[comps.C7](&world)
	c8ID := ecs.ComponentID[comps.C8](&world)

	extraIDs := []ecs.ID{c1ID, c2ID, c3ID, c4ID, c5ID, c6ID, c7ID, c8ID}

	world.Batch().New(n, posID, velID)

	ids := []ecs.ID{}
	for i := range n * 4 {
		ids = append(ids, posID)
		for j, id := range extraIDs {
			m := 1 << j
			if i&m == m {
				ids = append(ids, id)
			}
		}
		world.NewEntity(ids...)

		ids = ids[:0]
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
	query := world.Query(&filter)
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
	c1ID := ecs.ComponentID[comps.C1](&world)
	c2ID := ecs.ComponentID[comps.C2](&world)
	c3ID := ecs.ComponentID[comps.C3](&world)
	c4ID := ecs.ComponentID[comps.C4](&world)
	c5ID := ecs.ComponentID[comps.C5](&world)
	c6ID := ecs.ComponentID[comps.C6](&world)
	c7ID := ecs.ComponentID[comps.C7](&world)
	c8ID := ecs.ComponentID[comps.C8](&world)

	extraIDs := []ecs.ID{c1ID, c2ID, c3ID, c4ID, c5ID, c6ID, c7ID, c8ID}

	world.Batch().New(n, posID, velID)

	ids := []ecs.ID{}
	for i := range n * 4 {
		ids = append(ids, posID)
		for j, id := range extraIDs {
			m := 1 << j
			if i&m == m {
				ids = append(ids, id)
			}
		}
		world.NewEntity(ids...)

		ids = ids[:0]
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
	query := world.Query(&cf)
	for query.Next() {
		pos := (*comps.Position)(query.Get(posID))
		sum += pos.X + pos.Y
	}
	runtime.KeepAlive(sum)
}
