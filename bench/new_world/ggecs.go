package newworld

import (
	"testing"

	ecs "github.com/marioolofo/go-gameengine-ecs"
	"github.com/stretchr/testify/assert"
)

func runGGEcs(b *testing.B, n int) {
	var world ecs.World
	for b.Loop() {
		world = ecs.NewWorld(1024)
	}
	e := world.NewEntity()
	assert.True(b, world.IsAlive(e))
}
