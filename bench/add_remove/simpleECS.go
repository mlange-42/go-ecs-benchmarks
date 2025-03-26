package addremove

import (
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runSimpleECS(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.New(n)
	ecs.Register2[comps.Position, comps.Velocity](world)

	for range n {
		e := ecs.NewEntity(world)
		ecs.Add(world, e, comps.Position{})
	}
	b.StartTimer()
	for b.Loop() {
		entities := ecs.GetStorage[comps.Position](world).And(nil)
		for _, e := range entities {
			ecs.Add(world, e, comps.Velocity{})
		}
		for _, e := range entities {
			ecs.Remove[comps.Velocity](world, e)
		}
	}
}
