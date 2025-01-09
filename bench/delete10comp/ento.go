package delete10comp

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/wfranczyk/ento"
)

func runEnto(b *testing.B, n int) {
	b.StopTimer()
	world := ento.NewWorldBuilder().
		WithDenseComponents(comps.C1{}).
		WithDenseComponents(comps.C2{}).
		WithDenseComponents(comps.C3{}).
		WithDenseComponents(comps.C4{}).
		WithDenseComponents(comps.C5{}).
		WithDenseComponents(comps.C6{}).
		WithDenseComponents(comps.C7{}).
		WithDenseComponents(comps.C8{}).
		WithDenseComponents(comps.C9{}).
		WithDenseComponents(comps.C10{}).
		Build(1024)

	entities := make([]*ento.Entity, 0, n)
	for range n {
		e := world.AddEntity(
			comps.C1{}, comps.C2{}, comps.C3{}, comps.C4{}, comps.C5{},
			comps.C6{}, comps.C7{}, comps.C8{}, comps.C9{}, comps.C10{},
		)
		entities = append(entities, e)
	}

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		for _, e := range entities {
			world.RemoveEntity(e)
		}
		b.StopTimer()

		entities = entities[:0]
		for range n {
			e := world.AddEntity(
				comps.C1{}, comps.C2{}, comps.C3{}, comps.C4{}, comps.C5{},
				comps.C6{}, comps.C7{}, comps.C8{}, comps.C9{}, comps.C10{},
			)
			entities = append(entities, e)
		}
	}
}
