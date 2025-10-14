package query2comp

import (
	"fmt"
	"runtime"
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

	for i := 0; i < n*10; i++ {
		e := world.CreateEntity(strconv.Itoa(i))
		volt.AddComponent(world, e, comps.Position{})
	}
	for i := 0; i < n; i++ {
		e := world.CreateEntity(strconv.Itoa(i))
		volt.AddComponents2(world, e, comps.Position{}, comps.Velocity{X: 1, Y: 1})
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
	if sum != float64(n*b.N*2) {
		panic(fmt.Sprintf("Expected sum %d, got %.2f", n*b.N*2, sum))
	}
	runtime.KeepAlive(sum)
}
