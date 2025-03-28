package delete10comp

import (
	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"testing"
)

func runSimpleECS(b *testing.B, n int) {
	world := ecs.New(n)
	entities := make([]ecs.Entity, 0, n)
	for range n {
		e := ecs.NewEntity(world)
		ecs.Add3(world, e, comps.C1{}, comps.C2{}, comps.C3{})
		ecs.Add3(world, e, comps.C4{}, comps.C5{}, comps.C6{})
		ecs.Add3(world, e, comps.C7{}, comps.C8{}, comps.C9{})
		ecs.Add(world, e, comps.C10{})
		entities = append(entities, e)
	}
	for b.Loop() {
		for _, e := range entities {
			ecs.Kill(world, e)
		}
	}
}
