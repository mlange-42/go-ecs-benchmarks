package newworld

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unitoftime/ecs"
)

func runUot(b *testing.B, n int) {
	var world *ecs.World
	for b.Loop() {
		world = ecs.NewWorld()
	}
	assert.NotNil(b, world)
}
