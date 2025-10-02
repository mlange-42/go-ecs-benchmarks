package addremove

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/yohamta/donburi"
)

func runDonburi(b *testing.B, n int) {
	world := donburi.NewWorld()

	var position = donburi.NewComponentType[comps.Position]()
	var velocity = donburi.NewComponentType[comps.Velocity]()

	entities := make([]donburi.Entity, 0, n)
	for range n {
		entities = append(entities, world.Create(position))
	}
	for _, e := range entities {
		donburi.Add(world.Entry(e), position, &comps.Position{})
	}

	// Iterate once for more fairness
	for _, e := range entities {
		donburi.Add(world.Entry(e), velocity, &comps.Velocity{})
	}
	for _, e := range entities {
		donburi.Remove[comps.Velocity](world.Entry(e), velocity)
	}

	for b.Loop() {
		for _, e := range entities {
			donburi.Add(world.Entry(e), velocity, &comps.Velocity{})
		}
		for _, e := range entities {
			donburi.Remove[comps.Velocity](world.Entry(e), velocity)
		}
	}
}
