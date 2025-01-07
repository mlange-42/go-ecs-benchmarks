package addremove

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

func runDonburi(b *testing.B, n int) {
	b.StopTimer()
	world := donburi.NewWorld()

	var position = donburi.NewComponentType[comps.Position]()
	var velocity = donburi.NewComponentType[comps.Velocity]()

	for i := 0; i < n; i++ {
		world.Create(position)
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

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		queryPos.Each(world, func(entry *donburi.Entry) {
			entry.AddComponent(velocity)
		})
		queryPosVel.Each(world, func(entry *donburi.Entry) {
			entry.RemoveComponent(velocity)
		})
	}
}
