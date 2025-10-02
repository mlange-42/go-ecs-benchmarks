package addremove

import (
	"strconv"
	"testing"

	"github.com/akmonengine/volt"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

type voltConfig = volt.ComponentConfig[volt.ComponentInterface]

func runVolt(b *testing.B, n int) {
	world := volt.CreateWorld(1024)

	volt.RegisterComponent[comps.Position](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.Velocity](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})

	entities := make([]volt.EntityId, 0, n)
	for i := 0; i < n; i++ {
		e := world.CreateEntity(strconv.Itoa(i))
		volt.AddComponent(world, e, comps.Position{})
		entities = append(entities, e)
	}

	// Iterate once for more fairness
	for _, e := range entities {
		volt.AddComponent(world, e, comps.Velocity{})
	}
	for _, e := range entities {
		volt.RemoveComponent[comps.Velocity](world, e)
	}

	for b.Loop() {
		for _, e := range entities {
			volt.AddComponent(world, e, comps.Velocity{})
		}
		for _, e := range entities {
			volt.RemoveComponent[comps.Velocity](world, e)
		}
	}
}
