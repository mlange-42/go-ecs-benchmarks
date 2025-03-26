package addremovelarge

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
		ecs.Add3(world, e, comps.Position{}, comps.C1{}, comps.C2{})
		ecs.Add3(world, e, comps.C3{}, comps.C4{}, comps.C5{})
		ecs.Add3(world, e, comps.C6{}, comps.C7{}, comps.C8{})
		ecs.Add(world, e, comps.C9{})
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
