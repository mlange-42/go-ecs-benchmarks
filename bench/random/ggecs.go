package random

import (
	"log"
	"math/rand/v2"
	"testing"

	ecs "github.com/marioolofo/go-gameengine-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
)

const (
	PositionComponentID ecs.ComponentID = iota
	VelocityComponentID
)

func runGGEcs(b *testing.B, n int) {
	b.StopTimer()
	world := ecs.NewWorld(1024)
	world.Register(ecs.NewComponentRegistry[comps.Position](PositionComponentID))

	entities := make([]ecs.EntityID, 0, n)
	for i := 0; i < n; i++ {
		e := world.NewEntity(PositionComponentID)
		entities = append(entities, e)
	}
	rand.Shuffle(n, util.Swap(entities))

	b.StartTimer()
	sum := 0.0
	for i := 0; i < b.N; i++ {
		for _, e := range entities {
			pos := (*comps.Position)(world.Component(e, PositionComponentID))
			sum += pos.X
		}
	}
	b.StopTimer()
	if sum > 0 {
		log.Fatal("error")
	}
}
