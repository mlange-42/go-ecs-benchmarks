package random

import (
	"log"
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/akmonengine/volt"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/mlange-42/go-ecs-benchmarks/bench/util"
)

type voltConfig = volt.ComponentConfig[volt.ComponentInterface]

func runVolt(b *testing.B, n int) {
	world := volt.CreateWorld(1024)

	volt.RegisterComponent[comps.Position](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})

	entities := make([]volt.EntityId, 0, n)
	for i := 0; i < n; i++ {
		e := world.CreateEntity(strconv.Itoa(i))
		volt.AddComponent(world, e, comps.Position{})
		entities = append(entities, e)
	}
	rand.Shuffle(n, util.Swap(entities))

	sum := 0.0
	// Don't use b.Loop and callback, as we do not want to measure
	// the cost of calling the non-inlined callback.
	b.ResetTimer()
	for range b.N {
		for _, e := range entities {
			pos := volt.GetComponent[comps.Position](world, e)
			sum += pos.X
		}
	}
	b.StopTimer()
	if sum > 0 {
		log.Fatal("error")
	}
}
