package query32arch

import (
	"testing"

	ecs "github.com/marioolofo/go-gameengine-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

const (
	PositionComponentID ecs.ComponentID = iota
	VelocityComponentID
	C1ID
	C2ID
	C3ID
	C4ID
	C5ID
)

func runGGEcs(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld(1024)
	world.Register(ecs.NewComponentRegistry[comps.Position](PositionComponentID))
	world.Register(ecs.NewComponentRegistry[comps.Velocity](VelocityComponentID))

	world.Register(ecs.NewComponentRegistry[comps.C1](C1ID))
	world.Register(ecs.NewComponentRegistry[comps.C2](C2ID))
	world.Register(ecs.NewComponentRegistry[comps.C3](C3ID))
	world.Register(ecs.NewComponentRegistry[comps.C4](C4ID))
	world.Register(ecs.NewComponentRegistry[comps.C5](C5ID))

	extraIDs := []ecs.ComponentID{C1ID, C2ID, C3ID, C4ID, C5ID}

	ids := []ecs.ComponentID{}
	for i := 0; i < n; i++ {
		ids = append(ids, PositionComponentID, VelocityComponentID)
		for j, id := range extraIDs {
			m := 1 << j
			if i&m == m {
				ids = append(ids, id)
			}
		}
		_ = world.NewEntity(ids...)

		ids = ids[:0]
	}

	mask := ecs.MakeComponentMask(PositionComponentID, VelocityComponentID)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		query := world.Query(mask)
		for query.Next() {
			pos := (*comps.Position)(query.Component(PositionComponentID))
			vel := (*comps.Velocity)(query.Component(VelocityComponentID))
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
}
