package query1in10

import (
	"testing"

	ecs "github.com/marioolofo/go-gameengine-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

const (
	PositionComponentID ecs.ComponentID = iota
	VelocityComponentID
)

func runGGEcs(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld(1024)
	world.Register(ecs.NewComponentRegistry[comps.Position](PositionComponentID))
	world.Register(ecs.NewComponentRegistry[comps.Velocity](VelocityComponentID))

	for i := 0; i < n*9; i++ {
		_ = world.NewEntity(PositionComponentID)
	}
	for i := 0; i < n; i++ {
		_ = world.NewEntity(PositionComponentID, VelocityComponentID)
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
