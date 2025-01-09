package delete10comp

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/unitoftime/ecs"
)

func runUot(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld()

	allIDs := []ecs.Component{
		ecs.C(comps.C1{}),
		ecs.C(comps.C2{}),
		ecs.C(comps.C3{}),
		ecs.C(comps.C4{}),
		ecs.C(comps.C5{}),
		ecs.C(comps.C6{}),
		ecs.C(comps.C7{}),
		ecs.C(comps.C8{}),
		ecs.C(comps.C9{}),
		ecs.C(comps.C10{}),
	}

	entities := make([]ecs.Id, 0, n)
	for range n {
		id := world.NewId()
		ecs.Write(world, id, allIDs...)
		entities = append(entities, id)
	}

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		for _, e := range entities {
			ecs.Delete(world, e)
		}
		b.StopTimer()

		entities = entities[:0]
		for range n {
			id := world.NewId()
			ecs.Write(world, id, allIDs...)
			entities = append(entities, id)
		}
	}
}
