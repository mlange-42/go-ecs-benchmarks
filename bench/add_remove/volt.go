package addremove

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

	for i := 0; i < n; i++ {
		e := world.CreateEntity(strconv.Itoa(i))
		volt.AddComponent(world, e, comps.Position{})
	}

	posMask := volt.CreateQuery1[comps.Position](world, volt.QueryConfiguration{})
	posVelMask := volt.CreateQuery2[comps.Position, comps.Velocity](world, volt.QueryConfiguration{})

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
