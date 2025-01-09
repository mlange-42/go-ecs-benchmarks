package delete2comp

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/wfranczyk/ento"
)

func runEnto(b *testing.B, n int) {
	b.StopTimer()
	world := ento.NewWorldBuilder().
		WithDenseComponents(comps.Position{}).
		WithDenseComponents(comps.Velocity{}).
		Build(1024)

	entities := make([]*ento.Entity, 0, n)
	for range n {
		e := world.AddEntity(comps.Position{}, comps.Velocity{})
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
			e := world.AddEntity(comps.Position{}, comps.Velocity{})
			entities = append(entities, e)
		}
	}
}
