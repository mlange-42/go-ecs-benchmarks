package newworld

import (
	"testing"

	"github.com/mlange-42/ark/ecs"
	"github.com/stretchr/testify/assert"
)

func runArk(b *testing.B, n int) {
	var world ecs.World
	for i := 0; i < b.N; i++ {
		world = ecs.NewWorld()
	}
	b.StopTimer()
	assert.False(b, world.IsLocked())
}
