package create2compalloc

import (
	"testing"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArk(b *testing.B, n int) {
	for b.Loop() {
		b.StopTimer()
		world := ecs.NewWorld(1024)

		mapper := ecs.NewMap2[comps.Position, comps.Velocity](&world)

		b.StartTimer()
		for range n {
			mapper.NewEntityFn(nil)
		}
	}
}

func runArkBatched(b *testing.B, n int) {
	for b.Loop() {
		b.StopTimer()
		world := ecs.NewWorld(1024)

		mapper := ecs.NewMap2[comps.Position, comps.Velocity](&world)

		b.StartTimer()
		mapper.NewBatchFn(n, nil)
	}
}
