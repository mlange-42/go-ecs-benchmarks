package query32arch

import (
	"runtime"
	"testing"

	"github.com/akmonengine/volt"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

func runVolt(b *testing.B, n int) {
	world := volt.CreateWorld(1024)

	volt.RegisterComponent[comps.Position](world, &volt.ComponentConfig[comps.Position]{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.Velocity](world, &volt.ComponentConfig[comps.Velocity]{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.C1](world, &volt.ComponentConfig[comps.C1]{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.C2](world, &volt.ComponentConfig[comps.C2]{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.C3](world, &volt.ComponentConfig[comps.C3]{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.C4](world, &volt.ComponentConfig[comps.C4]{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.C5](world, &volt.ComponentConfig[comps.C5]{BuilderFn: func(component any, configuration any) {}})
	extraComps := []volt.ComponentId{comps.C1Id, comps.C2Id, comps.C3Id, comps.C4Id, comps.C5Id}

	var ids []volt.ComponentIdConf
	for i := 0; i < n; i++ {
		e := world.CreateEntity()

		ids = append(ids, volt.ComponentIdConf{ComponentId: comps.PositionId}, volt.ComponentIdConf{ComponentId: comps.VelocityId})
		for j, id := range extraComps {
			m := 1 << j
			if i&m == m {
				ids = append(ids, volt.ComponentIdConf{ComponentId: id})
			}
		}

		world.AddComponents(e, ids...)
		ids = ids[:0]
	}

	query := volt.CreateQuery2[comps.Position, comps.Velocity](world, volt.QueryConfiguration{})

	loop := func() {
		for result := range query.Foreach(nil) {
			pos := result.A
			vel := result.B
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
	for b.Loop() {
		loop()
	}

	sum := 0.0
	for result := range query.Foreach(nil) {
		pos := result.A
		sum += pos.X + pos.Y
	}
	runtime.KeepAlive(sum)
}
