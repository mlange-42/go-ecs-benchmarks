package newworld

import (
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/stretchr/testify/assert"
)

func runArche(b *testing.B, n int) {
	var world ecs.World
	for i := 0; i < b.N; i++ {
		world = ecs.NewWorld()
		_ = ecs.ComponentID[comps.Position](&world)
		_ = ecs.ComponentID[comps.Velocity](&world)
	}
	b.StopTimer()
	assert.False(b, world.IsLocked())
}
