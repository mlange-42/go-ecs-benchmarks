package create2comp

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.New(n)
	ecs.Register2[comps.Position,comps.Velocity](world)
	for b.Loop() {
		b.StartTimer()
		for range n {
			e := ecs.NewEntity(world)
			ecs.Add2(world, e,
				comps.Position{},
				comps.Velocity{X: 1, Y: 1},
			)
		}
		b.StopTimer()
		entities := ecs.GetStorage[comps.Position](world).
			And(ecs.GetStorage[comps.Velocity](world))
		ecs.Kill(world, entities...)
	}
}
