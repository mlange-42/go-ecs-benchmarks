package create2comp

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/wfranczyk/ento"
)

func runEnto(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		world := ento.NewWorldBuilder().
			WithDenseComponents(comps.Position{}).
			WithDenseComponents(comps.Velocity{}).
			Build(1024)

		b.StartTimer()
		for range n {
			world.AddEntity(comps.Position{}, comps.Velocity{})
		}
		b.StopTimer()
	}
}
