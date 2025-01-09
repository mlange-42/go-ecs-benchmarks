package create10comp

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/component"
)

func runDonburi(b *testing.B, n int) {
	b.StopTimer()

	allIDs := []component.IComponentType{
		donburi.NewComponentType[comps.C1](),
		donburi.NewComponentType[comps.C2](),
		donburi.NewComponentType[comps.C3](),
		donburi.NewComponentType[comps.C4](),
		donburi.NewComponentType[comps.C5](),
		donburi.NewComponentType[comps.C6](),
		donburi.NewComponentType[comps.C7](),
		donburi.NewComponentType[comps.C8](),
		donburi.NewComponentType[comps.C9](),
		donburi.NewComponentType[comps.C10](),
	}

	world := donburi.NewWorld()

	entities := make([]donburi.Entity, 0, n)
	for range n {
		e := world.Create(allIDs...)
		entities = append(entities, e)
	}
	for _, e := range entities {
		world.Remove(e)
	}
	entities = entities[:0]

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		for range n {
			e := world.Create(allIDs...)
			entities = append(entities, e)
		}
		b.StopTimer()
		for _, e := range entities {
			world.Remove(e)
		}
		entities = entities[:0]
	}
}
