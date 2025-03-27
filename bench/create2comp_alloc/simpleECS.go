package create2compalloc

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	for b.Loop() {
		b.StopTimer()
		world := ecs.New(1024)
		ecs.Register2[comps.Position, comps.Velocity](world)
		b.StartTimer()
		for range n {
			e := ecs.NewEntity(world)
			ecs.Add2(world, e,
				comps.Position{},
				comps.Velocity{X: 1, Y: 1},
			)
		}
	}
}
