package newworld

import (
	"testing"

	"github.com/akmonengine/volt"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/stretchr/testify/assert"
)

type voltConfig = volt.ComponentConfig[volt.ComponentInterface]

func runVolt(b *testing.B, n int) {
	var world ecs.World
	for i := 0; i < b.N; i++ {
		world := volt.CreateWorld(1024)
		volt.RegisterComponent[comps.Position](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})
		volt.RegisterComponent[comps.Velocity](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})
	}
	b.StopTimer()
	assert.False(b, world.IsLocked())
}
