package addremovelarge

import (
	"testing"

	ecs "github.com/marioolofo/go-gameengine-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

// Component IDs
const (
	PositionComponentID ecs.ComponentID = iota
	VelocityComponentID
	C1ID
	C2ID
	C3ID
	C4ID
	C5ID
	C6ID
	C7ID
	C8ID
	C9ID
	C10ID
)

func runGGEcs(b *testing.B, n int) {
	world := ecs.NewWorld(1024)
	world.Register(ecs.NewComponentRegistry[comps.Position](PositionComponentID))
	world.Register(ecs.NewComponentRegistry[comps.Velocity](VelocityComponentID))

	world.Register(ecs.NewComponentRegistry[comps.C1](C1ID))
	world.Register(ecs.NewComponentRegistry[comps.C2](C2ID))
	world.Register(ecs.NewComponentRegistry[comps.C3](C3ID))
	world.Register(ecs.NewComponentRegistry[comps.C4](C4ID))
	world.Register(ecs.NewComponentRegistry[comps.C5](C5ID))
	world.Register(ecs.NewComponentRegistry[comps.C6](C6ID))
	world.Register(ecs.NewComponentRegistry[comps.C7](C7ID))
	world.Register(ecs.NewComponentRegistry[comps.C8](C8ID))
	world.Register(ecs.NewComponentRegistry[comps.C9](C9ID))
	world.Register(ecs.NewComponentRegistry[comps.C10](C10ID))

	entities := make([]ecs.EntityID, 0, n)
	for i := 0; i < n; i++ {
		entities = append(entities, world.NewEntity(
			PositionComponentID,
			C1ID, C2ID, C3ID, C4ID, C5ID,
			C6ID, C7ID, C8ID, C9ID, C10ID,
		))
	}
	// Iterate once for more fairness
	for _, e := range entities {
		world.AddComponent(e, VelocityComponentID)
	}
	for _, e := range entities {
		world.RemComponent(e, VelocityComponentID)
	}

	for b.Loop() {
		for _, e := range entities {
			world.AddComponent(e, VelocityComponentID)
		}
		for _, e := range entities {
			world.RemComponent(e, VelocityComponentID)
		}
	}
}
