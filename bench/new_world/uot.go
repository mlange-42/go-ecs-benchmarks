package newworld

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unitoftime/ecs"
)

func runUot(b *testing.B, n int) {
	var world *ecs.World
	for i := 0; i < b.N; i++ {
		world = ecs.NewWorld()
	}
	b.StopTimer()
	assert.NotNil(b, world)
}
