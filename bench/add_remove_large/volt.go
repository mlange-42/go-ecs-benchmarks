package addremovelarge

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

	volt.RegisterComponent[comps.C1](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.C2](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.C3](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.C4](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.C5](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.C6](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.C7](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.C8](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.C9](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.C10](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})

	entities := make([]volt.EntityId, 0, n)
	for i := 0; i < n; i++ {
		e, err := volt.CreateEntityWithComponents8(world, strconv.Itoa(i), comps.Position{},
			comps.C1{}, comps.C2{}, comps.C3{}, comps.C4{},
			comps.C5{}, comps.C6{}, comps.C7{},
		)
		if err != nil {
			panic("Volt crashed")
		}
		volt.AddComponents3(world, e, comps.C8{}, comps.C9{}, comps.C10{})
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
