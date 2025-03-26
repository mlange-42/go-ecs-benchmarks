package random

import (
	"log"
	"math/rand/v2"
	"testing"

	ecs "github.com/BrownNPC/simple-ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
)

func runSimpleECS(b *testing.B, n int) {
	world := ecs.New(n)

	for range n {
		e := ecs.NewEntity(world)
		ecs.Add(world, e,
			comps.Position{},
		)
	}
	POSITION := ecs.GetStorage[comps.Position](world)
	entities := POSITION.And(nil)
	rand.Shuffle(n, util.Swap(entities))
	

	sum := 0.0
	for b.Loop() {
		for _, e := range entities {
			pos := POSITION.Get(e)
			sum += pos.X
		}
	}
	if sum > 0 {
		log.Fatal("error")
	}
}
