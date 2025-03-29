package delete2comp

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	world := ecs.New(n)
	entities := make([]ecs.Entity, 0, n)
	ecs.Register2[comps.Position, comps.Velocity](world)
	for range n {
		e := ecs.NewEntity(world)
		ecs.Add2(world, e,
			comps.Position{},
			comps.Velocity{X: 1, Y: 1},
		)
		entities = append(entities, e)
	}
	for b.Loop() {
		for _, e := range entities {
			ecs.Kill(world, e)
		}
	}
}
