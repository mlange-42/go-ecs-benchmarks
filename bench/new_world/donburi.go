package newworld

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yohamta/donburi"
)

func runDonburi(b *testing.B, n int) {
	var world donburi.World
	for b.Loop() {
		world = donburi.NewWorld()
	}
	assert.Equal(b, 0, world.Len())
}
