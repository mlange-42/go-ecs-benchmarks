package newworld

import (
	"testing"

	ecs "github.com/marioolofo/go-gameengine-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/stretchr/testify/assert"
)

const (
	PositionComponentID ecs.ComponentID = iota
	VelocityComponentID
)

func runGGEcs(b *testing.B, n int) {
	var world ecs.World
	for i := 0; i < b.N; i++ {
		world = ecs.NewWorld(1024)
		world.Register(ecs.NewComponentRegistry[comps.Position](PositionComponentID))
		world.Register(ecs.NewComponentRegistry[comps.Velocity](VelocityComponentID))
	}
	b.StopTimer()
	e := world.NewEntity()
	assert.True(b, world.IsAlive(e))
}
