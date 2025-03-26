package newworld

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/stretchr/testify/assert"
)

func runSimpleECS(b *testing.B, n int) {
	var world *ecs.Pool
	for b.Loop() {
		world = ecs.New(1024)
	}
	assert.NotNil(b, world)
}
