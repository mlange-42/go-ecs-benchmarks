package create2compalloc

import (
	"testing"

	"github.com/akmonengine/volt"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

type voltConfig = volt.ComponentConfig[volt.ComponentInterface]

func runVolt(b *testing.B, n int) {
	for b.Loop() {
		b.StopTimer()
		world := volt.CreateWorld(1024)

		volt.RegisterComponent[comps.Position](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})
		volt.RegisterComponent[comps.Velocity](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})

		b.StartTimer()
		for range n {
			_, err := volt.CreateEntityWithComponents2(world, comps.Position{}, comps.Velocity{})
			if err != nil {
				panic("Volt crashed")
			}
		}
	}
}
