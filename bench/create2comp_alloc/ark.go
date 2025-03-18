package create2compalloc

import (
	"testing"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArk(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		world := ecs.NewWorld(1024)

		mapper := ecs.NewMap2[comps.Position, comps.Velocity](&world)

		b.StartTimer()
		for range n {
			mapper.NewEntityFn(nil)
		}
		b.StopTimer()
	}
}

func runArkBatched(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		world := ecs.NewWorld(1024)

		mapper := ecs.NewMap2[comps.Position, comps.Velocity](&world)

		b.StartTimer()
		mapper.NewBatchFn(n, nil)
		b.StopTimer()
	}
}
