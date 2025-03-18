package query2comp

import (
	"testing"

	"github.com/mlange-42/ark/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runArk(b *testing.B, n int) {
	world := ecs.NewWorld(1024)

	mapper := ecs.NewMap2[comps.Position, comps.Velocity](&world)
	for range n {
		_ = mapper.NewEntity(&comps.Position{}, &comps.Velocity{X: 1, Y: 1})
	}

	filter := ecs.NewFilter2[comps.Position, comps.Velocity](&world)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		query := filter.Query()
		for query.Next() {
			pos, vel := query.Get()
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
	b.StopTimer()

	query := filter.Query()
	for query.Next() {
		pos, _ := query.Get()
		if pos.X == 0 || pos.Y == 0 {
			panic("assertion failed")
		}
	}
}
