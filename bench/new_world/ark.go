package newworld

import (
	"testing"

	"github.com/mlange-42/ark/ecs"
	"github.com/stretchr/testify/assert"
)

func runArk(b *testing.B, n int) {
	var world *ecs.World
	for b.Loop() {
		world = ecs.NewWorld()
	}
	assert.False(b, world.IsLocked())
}
