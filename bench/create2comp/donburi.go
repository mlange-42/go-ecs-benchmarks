package create2comp

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/component"
)

func runDonburi(b *testing.B, n int) {
	allIDs := []component.IComponentType{
		donburi.NewComponentType[comps.Position](),
		donburi.NewComponentType[comps.Velocity](),
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

	for b.Loop() {
		for range n {
			e := world.Create(allIDs...)
			entities = append(entities, e)
		}
		b.StopTimer()
		for _, e := range entities {
			world.Remove(e)
		}
		entities = entities[:0]
		b.StartTimer()
	}
}
