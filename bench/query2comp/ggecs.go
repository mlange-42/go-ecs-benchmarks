package query2comp

import (
	"testing"

	ecs "github.com/marioolofo/go-gameengine-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

const (
	positionComponentID ecs.ComponentID = iota
	velocityComponentID
)

func runGGEcs(b *testing.B, n int) {
	world := ecs.NewWorld(1024)
	world.Register(ecs.NewComponentRegistry[comps.Position](positionComponentID))
	world.Register(ecs.NewComponentRegistry[comps.Velocity](velocityComponentID))

	for i := 0; i < n; i++ {
		_ = world.NewEntity(positionComponentID, velocityComponentID)
	}

	mask := ecs.MakeComponentMask(positionComponentID, velocityComponentID)

	for b.Loop() {
		query := world.Query(mask)
		for query.Next() {
			pos := (*comps.Position)(query.Component(positionComponentID))
			vel := (*comps.Velocity)(query.Component(velocityComponentID))
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
}
