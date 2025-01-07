package create2comp

import (
	"testing"

	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
	"github.com/yohamta/donburi"
)

func runDonburi(b *testing.B, n int) {
	b.StopTimer()
	var position = donburi.NewComponentType[comps.Position]()
	var velocity = donburi.NewComponentType[comps.Velocity]()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		world := donburi.NewWorld()

		b.StartTimer()
		for range n {
			world.Create(position, velocity)
		}
		b.StopTimer()
	}
}
