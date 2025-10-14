package query256arch

import (
	"runtime"
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
	world.Register(ecs.NewComponentRegistry[comps.C5](C6ID))
	world.Register(ecs.NewComponentRegistry[comps.C5](C7ID))
	world.Register(ecs.NewComponentRegistry[comps.C5](C8ID))

	extraIDs := []ecs.ComponentID{C1ID, C2ID, C3ID, C4ID, C5ID, C6ID, C7ID, C8ID}

	for range n {
		_ = world.NewEntity(PositionComponentID, VelocityComponentID)
	}

	ids := []ecs.ComponentID{}
	for i := range n * 4 {
		ids = append(ids, PositionComponentID)
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

	loop := func() {
		query := world.Query(mask)
		for query.Next() {
			pos := (*comps.Position)(query.Component(PositionComponentID))
			vel := (*comps.Velocity)(query.Component(VelocityComponentID))
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
	for b.Loop() {
		loop()
	}

	sum := 0.0
	query := world.Query(mask)
	for query.Next() {
		pos := (*comps.Position)(query.Component(PositionComponentID))
		sum += pos.X + pos.Y
	}
	runtime.KeepAlive(sum)
}
