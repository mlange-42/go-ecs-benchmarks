package create2comp

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

	allIDs := []ecs.ComponentID{
		PositionComponentID,
		VelocityComponentID,
	}

	entities := make([]ecs.EntityID, 0, n)
	for range n {
		e := world.NewEntity(allIDs...)
		entities = append(entities, e)
	}
	for _, e := range entities {
		world.RemEntity(e)
	}
	entities = entities[:0]

	for b.Loop() {
		for range n {
			e := world.NewEntity(allIDs...)
			entities = append(entities, e)
		}
		b.StopTimer()
		for _, e := range entities {
			world.RemEntity(e)
		}
		entities = entities[:0]
		b.StartTimer()
	}
}
