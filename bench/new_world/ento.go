package newworld

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/stretchr/testify/assert"
	"github.com/wfranczyk/ento"
)

func runEnto(b *testing.B, n int) {
	var world *ento.World
	for i := 0; i < b.N; i++ {
		world = ento.NewWorldBuilder().
			WithDenseComponents(comps.Position{}).
			WithDenseComponents(comps.Velocity{}).
			Build(1024)
	}
	b.StopTimer()
	assert.NotNil(b, world)
}
