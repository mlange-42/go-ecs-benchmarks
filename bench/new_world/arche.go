package newworld

import (
	"testing"

	"github.com/mlange-42/arche/ecs"
	"github.com/stretchr/testify/assert"
)

func runArche(b *testing.B, n int) {
	var world ecs.World
	for b.Loop() {
		world = ecs.NewWorld()
	}
	assert.False(b, world.IsLocked())
}
