package create2compalloc

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
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		world := ecs.NewWorld(1024)
		world.Register(ecs.NewComponentRegistry[comps.Position](PositionComponentID))
		world.Register(ecs.NewComponentRegistry[comps.Velocity](VelocityComponentID))

		allIDs := []ecs.ComponentID{
			PositionComponentID,
			VelocityComponentID,
		}

		b.StartTimer()
		for range n {
			_ = world.NewEntity(allIDs...)
		}
		b.StopTimer()
	}
}
