package query2comp1of10

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

	for i := 0; i < n*9; i++ {
		world.Create(position)
	}
	for i := 0; i < n; i++ {
		world.Create(position, velocity)
	}

	query := donburi.NewQuery(filter.Contains(position, velocity))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		query.Each(world, func(entry *donburi.Entry) {
			pos := position.Get(entry)
			vel := velocity.Get(entry)

			pos.X += vel.X
			pos.Y += vel.Y
		})
	}
}
