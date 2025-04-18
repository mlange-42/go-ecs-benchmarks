package addremovelarge

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/component"
	"github.com/yohamta/donburi/filter"
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

	for i := 0; i < n; i++ {
		world.Create(allIDs...)
	}

	queryPos := donburi.NewQuery(filter.Contains(position))
	queryPosVel := donburi.NewQuery(filter.Contains(position, velocity))

	// Iterate once for more fairness
	queryPos.Each(world, func(entry *donburi.Entry) {
		entry.AddComponent(velocity)
	})
	queryPosVel.Each(world, func(entry *donburi.Entry) {
		entry.RemoveComponent(velocity)
	})

	for b.Loop() {
		queryPos.Each(world, func(entry *donburi.Entry) {
			entry.AddComponent(velocity)
		})
		queryPosVel.Each(world, func(entry *donburi.Entry) {
			entry.RemoveComponent(velocity)
		})
	}
}
