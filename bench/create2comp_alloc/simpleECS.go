package create2compalloc

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	world := ecs.New(1024).EnableGrowing()
	ecs.Register2[comps.Position, comps.Velocity](world)
	for b.Loop() {
		for range n {
			e := ecs.NewEntity(world)
			ecs.Add2(world, e,
				comps.Position{},
				comps.Velocity{X: 1, Y: 1},
			)
		}
	}
}
