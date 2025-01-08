package newworld

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/stretchr/testify/assert"
	"github.com/yohamta/donburi"
)

func runDonburi(b *testing.B, n int) {
	var world donburi.World
	for i := 0; i < b.N; i++ {
		world = donburi.NewWorld()
		_ = donburi.NewComponentType[comps.Position]()
		_ = donburi.NewComponentType[comps.Velocity]()
	}
	b.StopTimer()
	assert.Equal(b, 0, world.Len())
}
