package query2comp

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

func runDonburi(b *testing.B, n int) {
	world := donburi.NewWorld()

	var position = donburi.NewComponentType[comps.Position]()
	var velocity = donburi.NewComponentType[comps.Velocity]()

	for i := 0; i < n*10; i++ {
		world.Create(position)
	}
	for i := 0; i < n; i++ {
		e := world.Create(position, velocity)
		entry := world.Entry(e)
		vel := (*comps.Velocity)(entry.Component(velocity))
		vel.X = 1
		vel.Y = 1
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
	if sum != float64(n*b.N*2) {
		panic(fmt.Sprintf("Expected sum %d, got %.2f", n*b.N*2, sum))
	}
	runtime.KeepAlive(sum)
}
