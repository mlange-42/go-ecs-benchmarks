package newworld

import (
	"testing"

	"github.com/akmonengine/volt"
	"github.com/stretchr/testify/assert"
)

type voltConfig = volt.ComponentConfig[volt.ComponentInterface]

func runVolt(b *testing.B, n int) {
	var world *volt.World
	for b.Loop() {
		world = volt.CreateWorld(1024)
	}
	assert.NotNil(b, world)
}
