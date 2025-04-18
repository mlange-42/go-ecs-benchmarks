package addremove

import (
	"testing"

	ecs "github.com/marioolofo/go-gameengine-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

// Component IDs
const (
	PositionComponentID ecs.ComponentID = iota
	VelocityComponentID
)

func runGGEcs(b *testing.B, n int) {
	world := ecs.NewWorld(1024)
	world.Register(ecs.NewComponentRegistry[comps.Position](PositionComponentID))
	world.Register(ecs.NewComponentRegistry[comps.Velocity](VelocityComponentID))

	for i := 0; i < n; i++ {
		_ = world.NewEntity(PositionComponentID)
	}

	posMask := ecs.MakeComponentMask(PositionComponentID)
	posVelMask := ecs.MakeComponentMask(PositionComponentID, VelocityComponentID)

	entities := make([]ecs.EntityID, 0, n)

	// Iterate once for more fairness
	query := world.Query(posMask)
	for query.Next() {
		entities = append(entities, query.Entity())
	}

	for _, e := range entities {
		world.AddComponent(e, VelocityComponentID)
	}

	entities = entities[:0]
	query = world.Query(posVelMask)
	for query.Next() {
		entities = append(entities, query.Entity())
	}

	for _, e := range entities {
		world.RemComponent(e, VelocityComponentID)
	}
	entities = entities[:0]

	for b.Loop() {
		query := world.Query(posMask)
		for query.Next() {
			entities = append(entities, query.Entity())
		}

		for _, e := range entities {
			world.AddComponent(e, VelocityComponentID)
		}

		entities = entities[:0]
		query = world.Query(posVelMask)
		for query.Next() {
			entities = append(entities, query.Entity())
		}

		for _, e := range entities {
			world.RemComponent(e, VelocityComponentID)
		}
		entities = entities[:0]
	}
}
