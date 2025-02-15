package addremovelarge

import (
	"testing"

	"github.com/akmonengine/volt"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

type voltConfig = volt.ComponentConfig[volt.ComponentInterface]

func runVolt(b *testing.B, n int) {
	b.StopTimer()
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

	for i := 0; i < n; i++ {
		e, err := volt.CreateEntityWithComponents8(world, "-", comps.Position{},
			comps.C1{}, comps.C2{}, comps.C3{}, comps.C4{},
			comps.C5{}, comps.C6{}, comps.C7{},
		)
		if err != nil {
			panic("Volt crashed")
		}
		volt.AddComponents3(world, e, comps.C8{}, comps.C9{}, comps.C10{})
	}

	posMask := volt.CreateQuery1[comps.Position](world, []volt.OptionalComponent{})
	posVelMask := volt.CreateQuery2[comps.Position, comps.Velocity](world, []volt.OptionalComponent{})

	entities := make([]volt.EntityId, 0, n)

	// Iterate once for more fairness
	for result := range posMask.Foreach(nil) {
		entities = append(entities, result.EntityId)
	}

	for _, e := range entities {
		volt.AddComponent(world, e, comps.Velocity{})
	}

	entities = entities[:0]
	for result := range posVelMask.Foreach(nil) {
		entities = append(entities, result.EntityId)
	}

	for _, e := range entities {
		volt.RemoveComponent[comps.Velocity](world, e)
	}
	entities = entities[:0]

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for result := range posMask.Foreach(nil) {
			entities = append(entities, result.EntityId)
		}

		for _, e := range entities {
			volt.AddComponent(world, e, comps.Velocity{})
		}

		entities = entities[:0]
		for result := range posVelMask.Foreach(nil) {
			entities = append(entities, result.EntityId)
		}

		for _, e := range entities {
			volt.RemoveComponent[comps.Velocity](world, e)
		}
		entities = entities[:0]
	}
}
