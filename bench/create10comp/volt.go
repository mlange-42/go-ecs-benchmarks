package create10comp

import (
	"strconv"
	"testing"

	"github.com/akmonengine/volt"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

type voltConfig = volt.ComponentConfig[volt.ComponentInterface]

func runVolt(b *testing.B, n int) {
	b.StopTimer()
	world := volt.CreateWorld(1024)

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
	for id := range n {
		e, err := volt.CreateEntityWithComponents8(world, strconv.Itoa(id),
			comps.C1{}, comps.C2{}, comps.C3{}, comps.C4{},
			comps.C5{}, comps.C6{}, comps.C7{}, comps.C8{},
		)
		if err != nil {
			panic("Volt crashed")
		}
		volt.AddComponents2(world, e, comps.C9{}, comps.C10{})
		entities = append(entities, e)
	}
	for _, e := range entities {
		world.RemoveEntity(e)
	}
	entities = entities[:0]

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		for id := range n {
			e, err := volt.CreateEntityWithComponents8(world, strconv.Itoa(id),
				comps.C1{}, comps.C2{}, comps.C3{}, comps.C4{},
				comps.C5{}, comps.C6{}, comps.C7{}, comps.C8{},
			)
			if err != nil {
				panic("Volt crashed")
			}
			volt.AddComponents2(world, e, comps.C9{}, comps.C10{})
			entities = append(entities, e)
		}
		b.StopTimer()
		for _, e := range entities {
			world.RemoveEntity(e)
		}
		entities = entities[:0]
	}
}
