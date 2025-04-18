package delete2comp

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/yohamta/donburi"
)

func runDonburi(b *testing.B, n int) {
	var position = donburi.NewComponentType[comps.Position]()
	var velocity = donburi.NewComponentType[comps.Velocity]()

	world := donburi.NewWorld()

	entities := make([]donburi.Entity, 0, n)
	for range n {
		e := world.Create(position, velocity)
		entities = append(entities, e)
	}

	for b.Loop() {
		for _, e := range entities {
			world.Remove(e)
		}
		b.StopTimer()

		entities = entities[:0]
		for range n {
			e := world.Create(position, velocity)
			entities = append(entities, e)
		}
		b.StartTimer()
	}
}
