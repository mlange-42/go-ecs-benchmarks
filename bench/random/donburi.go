package random

import (
	"log"
	"math/rand/v2"
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
	"github.com/yohamta/donburi"
)

func runDonburi(b *testing.B, n int) {
	b.StopTimer()
	world := donburi.NewWorld()

	var position = donburi.NewComponentType[comps.Position]()

	entities := make([]donburi.Entity, 0, n)
	for i := 0; i < n; i++ {
		e := world.Create(position)
		entities = append(entities, e)
	}
	rand.Shuffle(n, util.Swap(entities))

	b.StartTimer()
	sum := 0.0
	for i := 0; i < b.N; i++ {
		for _, e := range entities {
			entry := world.Entry(e)
			pos := (*comps.Position)(entry.Component(position))
			sum += pos.X
		}
	}
	b.StopTimer()
	if sum > 0 {
		log.Fatal("error")
	}
}
