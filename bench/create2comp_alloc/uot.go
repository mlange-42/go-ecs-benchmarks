package create2compalloc

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/unitoftime/ecs"
)

func runUot(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		world := ecs.NewWorld()

		allIDs := []ecs.Component{
			ecs.C(comps.Position{}),
			ecs.C(comps.Velocity{}),
		}

		b.StartTimer()
		for range n {
			id := world.NewId()
			ecs.Write(world, id, allIDs...)
		}
		b.StopTimer()
	}
}
