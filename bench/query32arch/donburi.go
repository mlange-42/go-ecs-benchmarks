package query32arch

import (
	"runtime"
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
	var c1 = donburi.NewComponentType[comps.C1]()
	var c2 = donburi.NewComponentType[comps.C2]()
	var c3 = donburi.NewComponentType[comps.C3]()
	var c4 = donburi.NewComponentType[comps.C4]()
	var c5 = donburi.NewComponentType[comps.C5]()

	extraIDs := []component.IComponentType{c1, c2, c3, c4, c5}
	ids := []component.IComponentType{}

	for i := range n {
		ids = append(ids, position, velocity)
		for j, id := range extraIDs {
			m := 1 << j
			if i&m == m {
				ids = append(ids, id)
			}
		}

		world.Create(ids...)

		ids = ids[:0]
	}

	query := donburi.NewQuery(filter.Contains(position, velocity))

	loop := func() {
		query.Each(world, func(entry *donburi.Entry) {
			pos := position.Get(entry)
			vel := velocity.Get(entry)

			pos.X += vel.X
			pos.Y += vel.Y
		})
	}
	for b.Loop() {
		loop()
	}

	sum := 0.0
	query.Each(world, func(entry *donburi.Entry) {
		pos := position.Get(entry)
		sum += pos.X + pos.Y
	})
	runtime.KeepAlive(sum)
}
