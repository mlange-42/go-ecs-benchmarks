package create2comp

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

	volt.RegisterComponent[comps.Position](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.Velocity](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})

	entities := make([]volt.EntityId, 0, n)
	for id := range n {
		e, err := volt.CreateEntityWithComponents2(world, strconv.Itoa(id), comps.Position{}, comps.Velocity{})
		if err != nil {
			panic("Volt crashed")
		}
		entities = append(entities, e)
	}
	for _, e := range entities {
		world.RemoveEntity(e)
	}
	entities = entities[:0]

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		for id := range n {
			e, err := volt.CreateEntityWithComponents2(world, strconv.Itoa(id), comps.Position{}, comps.Velocity{})
			if err != nil {
				panic("Volt crashed")
			}
			entities = append(entities, e)
		}
		b.StopTimer()
		for _, e := range entities {
			world.RemoveEntity(e)
		}
		entities = entities[:0]
	}
}
