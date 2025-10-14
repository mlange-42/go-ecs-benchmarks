package random

import (
	"log"
	"math/rand/v2"
	"testing"

	ecs "github.com/marioolofo/go-gameengine-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
)

// Component IDs
const (
	positionComponentID ecs.ComponentID = iota
	velocityComponentID
)

func runGGEcs(b *testing.B, n int) {
	world := ecs.NewWorld(1024)
	world.Register(ecs.NewComponentRegistry[comps.Position](positionComponentID))

	entities := make([]ecs.EntityID, 0, n)
	for i := 0; i < n; i++ {
		e := world.NewEntity(positionComponentID)
		entities = append(entities, e)
	}
	rand.Shuffle(n, util.Swap(entities))

	sum := 0.0
	loop := func() {
		for _, e := range entities {
			pos := (*comps.Position)(world.Component(e, positionComponentID))
			sum += pos.X
		}
	}
	for b.Loop() {
		loop()
	}
	if sum > 0 {
		log.Fatal("error")
	}
}
