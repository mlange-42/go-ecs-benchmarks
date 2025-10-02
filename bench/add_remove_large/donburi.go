package addremovelarge

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/component"
)

func runDonburi(b *testing.B, n int) {
	world := donburi.NewWorld()

	var position = donburi.NewComponentType[comps.Position]()
	var velocity = donburi.NewComponentType[comps.Velocity]()

	allIDs := []component.IComponentType{
		position,
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

	entities := make([]donburi.Entity, 0, n)
	for range n {
		entities = append(entities, world.Create(allIDs...))
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
