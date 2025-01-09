package create2compalloc

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/component"
)

func runDonburi(b *testing.B, n int) {
	b.StopTimer()
	allIDs := []component.IComponentType{
		donburi.NewComponentType[comps.Position](),
		donburi.NewComponentType[comps.Velocity](),
	}

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		world := donburi.NewWorld()

		b.StartTimer()
		for range n {
			world.Create(allIDs...)
		}
		b.StopTimer()
	}
}
